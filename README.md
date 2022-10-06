# ipfs-gateway

A headless cli to run an ephemeral IPFS node with read-only local gateway, example use cases:
* Local retrieval of files distributed over IPFS
* Interface from a web2.0 applications to web3.0 over IPFS

## To run
1. Clone the repo
2. Build the executable with command `go build -ldflags "-s -w" main.go`

## To include as library
1. Clone the repo
2. Build the static library with command `go build -buildmode=c-archive -o main.a main.go`