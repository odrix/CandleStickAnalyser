package main

import (
	"fmt"

	"klintt.io/detect/handlers/detectdaily"
)

func main() {

	pairs := []string{"BTCBUSD", "ETHBUSD", "BNBBUSD", "SOLBUSD", "ADABUSD",
		"XRPBUSD", "DOTBUSD", "DOGEBUSD", "AVAXBUSD", "MATICBUSD",
		"SHIBBUSD", "LTCBUSD", "TRXBUSD",
		"ALGOBUSD", "LINKBUSD", "BCHBUSD", "OKBBUSD", "UNIBUSD",
		"AXSBUSD", "XLMBUSD", "ATOMBUSD", "NEARBUSD", "FTTBUSD",
		"VETBUSD", "EOSBUSD", "FILBUSD", "EGLDBUSD", "ETCBUSD",
		"SANDBUSD", "MANABUSD", "XTZBUSD", "GALABUSD", "THETABUSD"}
	//for i := 0; i < len(candlesDesc); i++ {
	//DetectOnManyPairToday(pairs, "")
	detectdaily.DetectOnManyPairHour(pairs)
	//}
	fmt.Printf("end.")
}
