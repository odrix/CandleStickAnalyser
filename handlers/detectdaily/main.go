package detectdaily

import (
	"fmt"

	"github.com/binance-exchange/go-binance"
	"klintt.io/detect/detector"
)

func main() {

	pairs := []string{"BTCBUSD", "ETHBUSD", "BNBBUSD", "SOLBUSD", "ADABUSD"}
	// pairs := []string{"BTCBUSD", "ETHBUSD", "BNBBUSD", "SOLBUSD", "ADABUSD",
	// 	"XRPBUSD", "DOTBUSD", "DOGEBUSD", "LUNABUSD", "AVAXBUSD",
	// 	"SHIBBUSD", "CROBUSD", "MATICBUSD", "LTCBUSD", "TRXBUSD",
	// 	"ALGOBUSD", "LINKBUSD", "BCHBUSD", "OKBBUSD", "UNIBUSD",
	// 	"AXSBUSD", "XLMBUSD", "ATOMBUSD", "NEARBUSD", "FTTBUSD",
	// 	"VETBUSD", "EOSBUSD", "FILBUSD", "EGLDBUSD", "ETCBUSD",
	// 	"SANDBUSD", "MANABUSD", "XTZBUSD", "GALABUSD", "THETABUSD"}
	//for i := 0; i < len(candlesDesc); i++ {
	//DetectOnManyPairToday(pairs, "")
	DetectOnManyPairHour(pairs)
	//}
	fmt.Printf("end.")
}

func DetectOnManyPairToday(pairs []string, notifyOnlyEmail string) {

	for j := 0; j < len(pairs); j++ {

		pair := pairs[j]
		fmt.Println(pair)
		p := detector.Detect(pair, binance.Day)

		if p.Type != "" {
			trace(p)
			if notifyOnlyEmail != "" {
				notifyOneEmail(p, notifyOnlyEmail)
			} else {
				notifyContacts(p)
			}
		}
	}
}

func DetectOnManyPairHour(pairs []string) {

	for j := 0; j < len(pairs); j++ {

		pair := pairs[j]
		fmt.Println(pair)
		p := detector.Detect(pair, binance.Day)

		if p.Type != "" {
			trace(p)
			notifyTwitter(p)
		}
	}
}

func trace(pattern detector.Pattern) {
	fmt.Printf("%s on %s \n", pattern.Type, pattern.Start)
}
