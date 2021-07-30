package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/wealdtech/go-ens"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/06318b3ca218411e87dcb18491711b56")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go resolve("ethereum.eth", client)
		wg.Add(1)
		go resolve("almonit.eth", client)
		wg.Add(1)
		go resolve("pepesza.eth", client)
		wg.Add(1)
		go resolve("alex.eth", client)
		wg.Add(1)
		go resolve("bitcoingenerator.eth", client)
	}
	wg.Wait()
}

var count = 0

func resolve(domain string, client *ethclient.Client) {
	start := time.Now()
	resol, _ := ens.NewResolver(client, domain) // most of the delay is here, and the first domain to be resolved gets the most delay
	q, _ := resol.Contenthash()
	CID, _ := ens.ContenthashToString(q)
	sh := shell.NewShell("localhost:5001")
	cat, _ := sh.Cat(CID)
	print(cat) // to avoid unused variable error
	elapsed := time.Since(start)
	log.Printf("time taken is %s", elapsed)
	count++
	print(count)
	wg.Done()
}
