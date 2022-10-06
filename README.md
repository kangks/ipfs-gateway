# ipfs-gateway

A headless cli to run an ephemeral IPFS node with read-only local gateway, example use cases:
* Local retrieval of files distributed over IPFS
* Interface from a web2.0 applications to web3.0 over IPFS

## To run
1. Clone the repo
2. Build for the target platform. List of supported platform can be found with `go tool dist list`.
    * Build the executable for OSX `GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ipfs-gateway main.go`
    * Rename the output to convention of `<filename>_<GOOS>_<GOARCH>`

## To include as library
1. Clone the repo
2. Build the static library with command `go build -buildmode=c-archive -o ipfs-gateway.a main.go`
    * Rename the output to convention of `<filename>_<GOOS>_<GOARCH>`