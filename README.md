# backrunner 
>* This repo contains a code for a simple DEX backrunner arbitrage bot. For simplicity this only looks for uniswaprouter02 trade.
>* This can be improved by adding more different trades


### Blue print 
- Design blueprint is attached [here](https://hackmd.io/@FdW5ADdtSn6Xozlgy3Rsyg/SJeDWSAz2) 


# Pre requisites
- Go 1.19
- Ubuntu 20.04 (any linux based distros) / OSX

# Build & Run
* Application can be build and started by using Makefile.
* Make sure to cd to project folder.
* Run the below commands in the terminal shell.
* Make sure to run Pre-run and Go path is set properly

# Pre-run
    make mod
    make lint

# How to run unit test
    make test

# How to run build
    make build

## How to run
* configure `.env.example`
``` text 
    export  ETH_CLIENT_URL - ETH websocket client 
    export FLASHBOTS_RELAY_URL - Flashbots bundle relay url"
    export BUNDLE_SINGING_KEY -  private key for signing flashbots bundle 
    export SENDER_SINGING_KEY - private key for signing transaction
```
* source your env file
* then start the server as `go run cmd/api/main.go`
