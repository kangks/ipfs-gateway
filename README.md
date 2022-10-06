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

# Simple performance comparison

To measure the key metrics such as TTFB, performance analyzer for Chrome extension can be used:
1. Install Performance Analyzer on [Chrome](https://chrome.google.com/webstore/detail/performance-analyser/djgfmlohefpomchfabngccpbaflcahjf) or on [Firefox](https://addons.mozilla.org/en-US/firefox/addon/performance-analyser/)
2. First load an example CID from (Cloudflare IPFS gateway)[https://cloudflare-ipfs.com/ipfs/QmbWqxBEKC3P8tqsKc98xmWNzrzDtRLMiMPL8wBuTGsMnR]
3. Once the page is loaded, runs the Performance Analyzer from the Chrome extension
4. Next load the same CID from the local gateway, such as http://localhost:8001/ipfs/QmbWqxBEKC3P8tqsKc98xmWNzrzDtRLMiMPL8wBuTGsMnR (assuming the gateway is serving on default port 8001)