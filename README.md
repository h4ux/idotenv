```
██ ██████   ██████  ████████ ███████ ███    ██ ██    ██ 
██ ██   ██ ██    ██    ██    ██      ████   ██ ██    ██ 
██ ██   ██ ██    ██    ██    █████   ██ ██  ██ ██    ██ 
██ ██   ██ ██    ██    ██    ██      ██  ██ ██  ██  ██  
██ ██████   ██████     ██    ███████ ██   ████   ████   
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
        set env variables and run command after double dash Ex. idotenv -run -- pnpm run dev
        Tip: you can also run: idotenv -- pnpm run dev
  -set string
        Set key value secret Ex: idotenv -set=KEY=VALUE (to add additional env var)
  -get string
        Get value of specific key Ex: idotenv -get=KEY
  -v    idotenv version
```

## How To Use

Add repository:

```
brew tap h4ux/idotenv
```

Install [idotenv](https://github.com/h4ux/idotenv):

```
brew install idotenv
```

Upgrade the idotenv CLI to the latest version:

```
brew upgrade idotenv
```

### Help

** Supports Mac OS (Intel, M1/2), Linux OS, Windows
