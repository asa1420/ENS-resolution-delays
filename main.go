package main

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/ethereum/go-ethereum/ethclient"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/wealdtech/go-ens"
	"log"
	"time"
)

func main() {
	f, err := excelize.OpenFile("measurement.xlsx")
	if err != nil {
		panic(err)
	}
	var domains = [5]string{"ethereum.eth", "almonit.eth", "pepesza.eth", "alex.eth", "bitcoingenerator.eth"}
	var cells = [5]string{"P5", "Q5", "R5", "S5", "T5"}
	for x, i := range domains {
		start := time.Now()
		client, err := ethclient.Dial("//./pipe/geth.ipc")
		if err != nil {
			panic(err)
		}
		resol, err := ens.NewResolver(client, i) // most of the delay is here, and the first domain to be resolved gets the most delay
		q, err := resol.Contenthash()
		CID, err := ens.ContenthashToString(q)
		sh := shell.NewShell("localhost:5001")
		cat, err := sh.Cat(CID)
		elapsed := time.Since(start)
		log.Printf("time taken is %s", elapsed)
		delay := elapsed.Seconds()
		f.SetCellFloat("Sheet1", cells[x], delay, 4, 64)
		print(cat) // to avoid unused variable error
	}
	f.SetCellValue("Sheet1", "O5", time.Now())
	err = f.Save()
	if err != nil {
		panic(err)
	}
}
