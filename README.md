# backrunner 
>* This repo containing code for a simple DEX backrunner arbitrage bot. For simplicity this only looks for uniswaprouter02 trade.
>* This can be improved by adding more different trades


### Blue print 
- Design blueprint is attached [here](https://hackmd.io/@FdW5ADdtSn6Xozlgy3Rsyg/SJeDWSAz2) 

### Architecture diagram
![plot](./backrunner_architecture.png)

# Pre requisites
- Go 1.19
- Ubuntu 20.04 (any linux based distros) / OSX

# Build & Run
* Application can be built and initiated by using Makefile.
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
    ETH_CLIENT_WSS_URL - ETH websocket client 
    FLASHBOTS_RELAY_URL - Flashbots bundle relay url
    BUNDLE_SINGING_KEY -  private key for signing flashbots bundle 
    SENDER_SINGING_KEY - private key for signing transaction
```
* source your env file
* then start the server as `go run cmd/api/main.go`
## Run as docker container
    make docker-image
    make docker-run
* Note:Make sure to update the .env.example file before running as docker container

> **Warning**
> This is just a protype. Running a mempool backrunner can be risky. There are poison tokens that drain your liquidity when you trade them. See [salmonella](https://github.com/Defi-Cartel/salmonella) for an example.
