```
   o            o                 o                                         
 _<|>_         <|>               <|>                                        
               < \               < >                                        
   o      o__ __o/    o__ __o     |        o__  __o   \o__ __o    o      o  
  <|>    /v     |    /v     v\    o__/_   /v      |>   |     |>  <|>    <|> 
  / \   />     / \  />       <\   |      />      //   / \   / \  < >    < > 
  \o/   \      \o/  \         /   |      \o    o/     \o/   \o/   \o    o/  
   |     o      |    o       o    o       v\  /v __o   |     |     v\  /v   
  / \    <\__  / \   <\__ __/>    <\__     <\/> __/>  / \   / \     <\/>    
``` 


![Release](https://github.com/h4ux/idotenv/actions/workflows/release.yml/badge.svg)

idotenv is a cli tool for injecting env variables from fetched url

Ex:

```
  -configure
        set the url to fetch from
  -d    idotenv debug (verbose)
  -list string {json, table, yaml}
        List key value secrets
  -run
        set env variables and run command after double dash Ex. idotenv run -- npm run dev
  -set string
        Set key value secret Ex: idotenv -set=KEY=VALUE (to add additional env var)
  -get string
        Get value of specific key Ex: idotenv -get=KEY
  -v    idotenv version
```

## Installation via install.sh

```bash
# binary will be in $(go env GOPATH)/bin/idotenv
curl -sSfL https://raw.githubusercontent.com/h4ux/idotenv/main/install.sh | sh -s -- -b $(go env GOPATH)/bin

# defualt installation into ./bin/
curl -sSfL https://raw.githubusercontent.com/h4ux/idotenv/main/install.sh | sh -s

```

Once you run idotenv -configure idotenv will create .env.idotenv file in executed path

### Help

** Supports Mac OS (Intel, M1/2), Linux OS, Windows
