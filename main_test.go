package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"

	icore "github.com/ipfs/interface-go-ipfs-core"
	icorepath "github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/ipfs/kubo/core"
)

var (
	//lint:ignore U1000 used only by skipped tests at present
	bigBuckBunnyCID = "QmbGtJg23skhvFmu9mJiePVByhfzu5rwo74MEkVDYAmF5T"
	imageCID        = "QmbWqxBEKC3P8tqsKc98xmWNzrzDtRLMiMPL8wBuTGsMnR"
	helloWorldCID   = "bafybeic35nent64fowmiohupnwnkfm2uxh6vpnyjlt3selcodjipfrokgi"
)

func newTestIPFSNode(t *testing.T, ctx context.Context) (icore.CoreAPI, *core.IpfsNode, error) {
	coreApi, ipfsNode, err := spawnAndBoostrap(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to spawn and connect to boostrap nodes: %s", err))
	}

	return coreApi, ipfsNode, err
}

func TestCidFetch(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	coreApi, _, err := newTestIPFSNode(t, ctx)
	if err != nil {
		panic(fmt.Errorf("failed creating a connected node: %s", err))
	}

	testPath := icorepath.New(testCID)
	obj, err := coreApi.Unixfs().Get(ctx, testPath)
	if err != nil {
		panic(fmt.Errorf("could not get file with CID: %s", err))
	}

	size, _ := obj.Size()
	if size <= 0 {
		t.Fatalf("expected object size > 0, actual %d", size)
	}
}

func TestGatewayFetch(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, ipfsNode, err := newTestIPFSNode(t, ctx)
	if err != nil {
		panic(fmt.Errorf("failed creating a connected node: %s", err))
	}

	go func() {
		if err = initGateway(ipfsNode); err != nil {
			panic(fmt.Errorf("failed serving gateway: %s", err))
		}
	}()

	time.Sleep(100 * time.Millisecond)

	res, err := http.Get("http://127.0.0.1:5001/ipfs/" + helloWorldCID)
	if err != nil {
		panic(fmt.Errorf("failed to connect to gateway: %s", err))
	}
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(fmt.Errorf("failed to read response body: %s", err))
	}
	log.Printf(string(out))

}
