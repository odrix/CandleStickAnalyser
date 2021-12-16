package detectdaily

import (
	"context"
	"fmt"
	"os"

	"github.com/binance-exchange/go-binance"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
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
		kl := getKline(logger, pair, binance.Day, 10)

		var candlesDesc []Candle
		for i := len(kl) - 1; i > 0; i-- {
			candlesDesc = append(candlesDesc, Candle{kl[i]})
		}

		yesterday := 1
		p := detectPattern(candlesDesc, yesterday, pair)

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

func detectPattern(candlesDesc []Candle, startDayIndex int, pair string) Pattern {
	p := Pattern{}
	if IsMorningStar(candlesDesc[startDayIndex:]) {
		p = Pattern{
			Pair:  pair,
			Type:  "Morning Star",
			Start: candlesDesc[startDayIndex+2].OpenTime,
			End:   candlesDesc[startDayIndex].CloseTime,
		}
	} else if IsEveningStar(candlesDesc[startDayIndex:]) {
		p = Pattern{
			Pair:  pair,
			Type:  "Evening Star",
			Start: candlesDesc[startDayIndex+2].OpenTime,
			End:   candlesDesc[startDayIndex].CloseTime,
		}
	} else if IsThreeWhiteSoldiers(candlesDesc[startDayIndex:]) {
		p = Pattern{
			Pair:  pair,
			Type:  "Three White Soldiers",
			Start: candlesDesc[startDayIndex+2].OpenTime,
			End:   candlesDesc[startDayIndex].CloseTime,
		}
	} else if IsThreeBlackCrows(candlesDesc[startDayIndex:]) {
		p = Pattern{
			Pair:  pair,
			Type:  "Three Black Crows",
			Start: candlesDesc[startDayIndex+2].OpenTime,
			End:   candlesDesc[startDayIndex].CloseTime,
		}
	} else if IsWhiteMarubozu(candlesDesc[startDayIndex:]) {
		p = Pattern{
			Pair:  pair,
			Type:  "White Marubozu",
			Start: candlesDesc[startDayIndex].OpenTime,
			End:   candlesDesc[startDayIndex].CloseTime,
		}
	} else if IsBlackMarubozu(candlesDesc[startDayIndex:]) {
		p = Pattern{
			Pair:  pair,
			Type:  "Black Marubozu",
			Start: candlesDesc[startDayIndex].OpenTime,
			End:   candlesDesc[startDayIndex].CloseTime,
		}
	} else if IsDoji(candlesDesc[startDayIndex:]) {
		p = Pattern{
			Pair:  pair,
			Type:  "Doji",
			Start: candlesDesc[startDayIndex].OpenTime,
			End:   candlesDesc[startDayIndex].CloseTime,
		}
	}
	return p
}

func trace(pattern Pattern) {
	fmt.Printf("%s on %s \n", pattern.Type, pattern.Start)
}

func getKline(logger log.Logger, symbol string, interval binance.Interval, limit int) []*binance.Kline {
	var ctx, _ = context.WithCancel(context.Background())

	binanceService := binance.NewAPIService(
		"https://www.binance.com",
		"", //os.Getenv("BINANCE_APIKEY"),
		nil,
		logger,
		ctx,
	)
	b := binance.NewBinance(binanceService)

	kl, err := b.Klines(binance.KlinesRequest{
		Symbol:   symbol,
		Interval: interval,
		Limit:    limit,
	})
	if err != nil {
		fmt.Printf("error on %s : %s \n", symbol, err.Error())
		var empty []*binance.Kline
		return empty
		// panic(err)
	}
	return kl
}
