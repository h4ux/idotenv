package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/joho/godotenv"
)

func fetchFile() []byte {
	url := goDotEnvVariable("URI_ADDR", ".env.idotenv")
	method := goDotEnvVariable("FETCH_METHOD", ".env.idotenv")
	data := goDotEnvVariable("FETCH_DATA", ".env.idotenv")

	client := &http.Client{}
	var dataReader = strings.NewReader(data)
	req, err := http.NewRequest(method, url, dataReader)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)
	return bodyText
}

func mergeMaps(m1 map[string]string, m2 map[string]string) {
	for k, v := range m2 {
		m1[k] = v
	}
}

func goDotEnvVariable(key string, envfile string) string {
	// load .env file
	//err := godotenv.Load(".env")

	//if err != nil {
	//home := os.Getenv("HOME")
	//err := godotenv.Load(envfile)
	myVars, err := godotenv.Read(envfile)

	if err != nil {
		log.Fatalf("Error loading " + envfile + " file")
	}
	//}
	if myVars[key] == "" {
		println(color.Colorize(color.Yellow, "variable "+key+" does not exists"))
	}

	return myVars[key]

	//if os.Getenv(key) == "" {
	//	//log.Fatalf("Error " + key + " not set in .env file")
	//	println(color.Colorize(color.Yellow, "variable "+key+" does not exists"))
	//}
	//return os.Getenv(key)
}

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

func isCommandAvailable(name string) bool {
	cmd := exec.Command("command", "-v", name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func createFile(path string, name string, content string) {
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString(content)
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		return text
	}
}

func configureVaultEnv(debug bool) {
	//home := os.Getenv("HOME")
	println(color.Colorize(color.Blue, "Please enter the uri address:"))
	uri_addr := getInput()
	println(color.Colorize(color.Blue, "Please enter the fetch method (POST, GET...):"))
	fetch_method := getInput()
	println(color.Colorize(color.Blue, "Please enter the fetch data:"))
	fetch_data := getInput()

	createFile("", ".env.idotenv", "URI_ADDR="+uri_addr+"\nFETCH_METHOD="+fetch_method+"\nFETCH_DATA="+fetch_data+"\n")
	createFile("", ".env.inject", "")
}

func list(debug bool) {
	//outputOptions := [4]string{"table", "json", "yaml", "pretty"}

	//if !itemExists(outputOptions, outputType) {
	//	outputType = "json"
	//}

	myEnv, err := godotenv.Read(".env.inject")
	if err != nil {
		log.Fatalf("Error loading .env.inject file")
	}

	filedata := fetchFile()

	output := string(filedata)
	fmt.Println(output)

	println(color.Colorize(color.Yellow, "This values will be added and override external vars if exists"))

	for key, value := range myEnv {
		println(color.Colorize(color.Cyan, key+": "+value))
	}
}

func set(debug bool, key string, val string) {
	envkey := goDotEnvVariable(key, ".env.inject")
	if envkey != "" {
		var approve string
		println(color.Colorize(color.Blue, "Looks like key: "+key+" exists, to Cancel press C or Enter to continue: "))
		fmt.Scanln(&approve)

		if strings.ToLower(approve) == "c" {
			os.Exit(1)
		}
	}

	env, err := godotenv.Unmarshal(key + "=" + val)

	if err != nil {
		log.Fatalf("Error creating env var")
		return
	}
	myEnv, err := godotenv.Read(".env.inject")

	if err != nil {
		log.Fatalf("Error reading .env.inject file")
		return
	}
	mergeMaps(myEnv, env)

	err = godotenv.Write(myEnv, ".env.inject")

	if err != nil {
		log.Fatalf("Error writing to .env.inject file")
	}

	if debug {
		fmt.Println(env)
	}
}

func run(debug bool) {
	filedata := fetchFile()
	reader := bytes.NewReader(filedata)

	myEnv, err := godotenv.Parse(reader)
	if err != nil {
		log.Fatalf("Error parsing .env string")
	}

	myInjectEnv, err := godotenv.Read(".env.inject")
	if err != nil {
		log.Fatalf("Error parsing .env.inject file")
	}

	//os.Clearenv()

	for key, value := range myEnv {
		if myInjectEnv[key] != "" {
			value = myInjectEnv[key]
		}
		os.Setenv(key, value)
		if debug {
			fmt.Println("Key:", key, "=>", "Value:", value)
		}
	}

	cmd := strings.Join(strings.Split(strings.Join(os.Args[1:], " "), "--")[1:], "--")
	if debug {
		fmt.Println(cmd)
	}

	parts := strings.Fields(cmd)
	datacmd := exec.Command(parts[0], parts[1:]...)
	datacmd.Stdout = os.Stdout
	datacmd.Stdin = os.Stdin
	datacmd.Stderr = os.Stderr
	datacmd.Run()
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func version() {
	println(color.Colorize(color.Cyan, `   

   o            o                 o                                         
 _<|>_         <|>               <|>                                        
               < \               < >                                        
   o      o__ __o/    o__ __o     |        o__  __o   \o__ __o    o      o  
  <|>    /v     |    /v     v\    o__/_   /v      |>   |     |>  <|>    <|> 
  / \   />     / \  />       <\   |      />      //   / \   / \  < >    < > 
  \o/   \      \o/  \         /   |      \o    o/     \o/   \o/   \o    o/  
   |     o      |    o       o    o       v\  /v __o   |     |     v\  /v   
  / \    <\__  / \   <\__ __/>    <\__     <\/> __/>  / \   / \     <\/>    idotenvVTAG
	`))
}

func main() {

	idotenvV := flag.Bool("v", false, "idotenv version")
	giDebug := flag.Bool("d", false, "idotenv debug (verbose)")
	giRun := flag.Bool("run", false, "set env variables and run command after double dash Ex. idotenv run -- npm run dev")
	giSet := flag.String("set", "", "Set key value secret Ex: idotenv -set=KEY=VALUE")
	giList := flag.Bool("list", false, "List key value secrets")
	giConfigure := flag.Bool("configure", false, "create configuration for the vault to be used")

	flag.Parse()

	if *idotenvV {
		version()
		return
	}

	if *giConfigure {
		configureVaultEnv(*giDebug)
		return
	}

	if *giSet != "" {
		s := strings.Split(*giSet, "=")
		if len(s) < 2 {
			println(color.Colorize(color.Yellow, "Wrong arguments received, Please refer to the help for the right usage"))
			return
		}

		key := s[0]
		value := s[1]
		set(*giDebug, key, value)
		return
	}

	if *giList {
		list(*giDebug)
		return
	}

	if *giRun {
		run(*giDebug)
		return
	}

	if *giDebug {
		fmt.Println("debug!!")
	}

	run(*giDebug)
}
