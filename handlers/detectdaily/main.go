package detectdaily

import (
	"fmt"
	"os"

	"github.com/binance-exchange/go-binance"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"klintt.io/detect/detector"
)

//const baseUrl string = "https://api.binance.com"

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = log.With(logger, "time", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	pairs := []string{"BTCUSDT", "ETHUSDT", "BNBUSDT", "SOLUSDT", "ADAUSDT"}
	// pairs := []string{"BTCUSDT", "ETHUSDT", "BNBUSDT", "SOLUSDT", "ADAUSDT",
	// 	"XRPUSDT", "DOTUSDT", "DOGEUSDT", "LUNAUSDT", "AVAXUSDT",
	// 	"SHIBUSDT", "CROUSDT", "MATICUSDT", "LTCUSDT", "TRXUSDT",
	// 	"ALGOUSDT", "LINKUSDT", "BCHUSDT", "OKBUSDT", "UNIUSDT",
	// 	"AXSUSDT", "XLMUSDT", "ATOMUSDT", "NEARUSDT", "FTTUSDT",
	// 	"VETUSDT", "EOSUSDT", "FILUSDT", "EGLDUSDT", "ETCUSDT",
	// 	"SANDUSDT", "MANAUSDT", "XTZUSDT", "GALAUSDT", "THETAUSDT"}
	//for i := 0; i < len(candlesDesc); i++ {
	DetectOnManyPairToday(pairs, "", logger)
	//}
	fmt.Printf("end.")
}

func DetectOnManyPairToday(pairs []string, notifyOnlyEmail string, logger log.Logger) {

	for j := 0; j < len(pairs); j++ {

		pair := pairs[j]
		fmt.Println(pair)
		kl := detector.GetKline(logger, pair, binance.Day, 10)

		var candlesDesc []detector.Candle
		for i := len(kl) - 1; i > 0; i-- {
			candlesDesc = append(candlesDesc, detector.KlineToCandle(kl[i]))
		}

		yesterday := 1
		p := detector.DetectPattern(candlesDesc, yesterday, pair)

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

func trace(pattern detector.Pattern) {
	fmt.Printf("%s on %s \n", pattern.Type, pattern.Start)
}
