package main

import (
	"fmt"

	"github.com/binance-exchange/go-binance"
	"klintt.io/detect/handlers/detectdaily"
)

func main() {

	pairs := []string{"BTCBUSD", "ETHBUSD", "BNBBUSD", "SOLBUSD", "ADABUSD",
		"XRPBUSD", "DOTBUSD", "DOGEBUSD", "AVAXBUSD", "MATICBUSD",
		"SHIBBUSD", "LTCBUSD", "TRXBUSD",
		"ALGOBUSD", "LINKBUSD", "BCHBUSD", "UNIBUSD",
		"AXSBUSD", "XLMBUSD", "ATOMBUSD", "NEARBUSD", "FTTBUSD",
		"VETBUSD", "EOSBUSD", "FILBUSD", "EGLDBUSD", "ETCBUSD",
		"SANDBUSD", "MANABUSD", "XTZBUSD", "GALABUSD", "THETABUSD"}
	detectdaily.DetectAndTweet(pairs, binance.Hour)

	fmt.Printf("end.")
}
