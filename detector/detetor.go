package detector

import (
	"context"
	"fmt"
	"os"

	"github.com/binance-exchange/go-binance"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
)

//const baseUrl string = "https://api.binance.com"

func detectPattern(candlesDesc []Candle, startCandleIndex int, pair string) Pattern {
	p := Pattern{}
	if IsMorningStar(candlesDesc[startCandleIndex:]) {
		p = Pattern{
			Pair:           pair,
			Type:           "Morning Star",
			TrendDirection: "Bullish",
			Start:          candlesDesc[startCandleIndex+2].OpenTime,
			End:            candlesDesc[startCandleIndex].CloseTime,
		}
	} else if IsEveningStar(candlesDesc[startCandleIndex:]) {
		p = Pattern{
			Pair:           pair,
			Type:           "Evening Star",
			TrendDirection: "Bearish",
			Start:          candlesDesc[startCandleIndex+2].OpenTime,
			End:            candlesDesc[startCandleIndex].CloseTime,
		}
	} else if IsThreeWhiteSoldiers(candlesDesc[startCandleIndex:]) {
		p = Pattern{
			Pair:           pair,
			Type:           "Three White Soldiers",
			TrendDirection: "Bullish",
			Start:          candlesDesc[startCandleIndex+2].OpenTime,
			End:            candlesDesc[startCandleIndex].CloseTime,
		}
	} else if IsThreeBlackCrows(candlesDesc[startCandleIndex:]) {
		p = Pattern{
			Pair:           pair,
			Type:           "Three Black Crows",
			TrendDirection: "Bearish",
			Start:          candlesDesc[startCandleIndex+2].OpenTime,
			End:            candlesDesc[startCandleIndex].CloseTime,
		}
	} else if IsWhiteMarubozu(candlesDesc[startCandleIndex:]) {
		p = Pattern{
			Pair:           pair,
			Type:           "White Marubozu",
			TrendDirection: "Bullish",
			Start:          candlesDesc[startCandleIndex].OpenTime,
			End:            candlesDesc[startCandleIndex].CloseTime,
		}
	} else if IsBlackMarubozu(candlesDesc[startCandleIndex:]) {
		p = Pattern{
			Pair:           pair,
			Type:           "Black Marubozu",
			TrendDirection: "Bearish",
			Start:          candlesDesc[startCandleIndex].OpenTime,
			End:            candlesDesc[startCandleIndex].CloseTime,
		}
	} else if IsHammer(candlesDesc[startCandleIndex:]) {
		p = Pattern{
			Pair:           pair,
			Type:           "Hammer",
			TrendDirection: "Bullish",
			Start:          candlesDesc[startCandleIndex].OpenTime,
			End:            candlesDesc[startCandleIndex].CloseTime,
		}
	} else if IsInvertedHammer(candlesDesc[startCandleIndex:]) {
		p = Pattern{
			Pair:           pair,
			Type:           "Inverted Hammer",
			TrendDirection: "Bearish",
			Start:          candlesDesc[startCandleIndex].OpenTime,
			End:            candlesDesc[startCandleIndex].CloseTime,
		}
	} else if IsDoji(candlesDesc[startCandleIndex:]) {
		p = Pattern{
			Pair:           pair,
			Type:           "Doji",
			TrendDirection: "Continuation",
			Start:          candlesDesc[startCandleIndex].OpenTime,
			End:            candlesDesc[startCandleIndex].CloseTime,
		}
	}
	return p
}

func getKline(logger log.Logger, symbol string, interval binance.Interval, limit int) []*binance.Kline {
	var ctx, cancel = context.WithCancel(context.Background())

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
	cancel()
	if err != nil {
		fmt.Printf("error on %s : %s \n", symbol, err.Error())
		var empty []*binance.Kline
		return empty
		// panic(err)
	}
	return kl
}

func klineToCandle(k *binance.Kline) Candle {
	return Candle{
		OpenTime:                 k.OpenTime,
		Open:                     k.Open,
		High:                     k.High,
		Low:                      k.Low,
		Close:                    k.Close,
		Volume:                   k.Volume,
		CloseTime:                k.CloseTime,
		QuoteAssetVolume:         k.QuoteAssetVolume,
		NumberOfTrades:           k.NumberOfTrades,
		TakerBuyBaseAssetVolume:  k.TakerBuyBaseAssetVolume,
		TakerBuyQuoteAssetVolume: k.TakerBuyQuoteAssetVolume}
}

func Detect(pair string, interval binance.Interval) Pattern {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = log.With(logger, "time", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	kl := getKline(logger, pair, interval, 10)

	var candlesDesc []Candle
	for i := len(kl) - 1; i > 0; i-- {
		candlesDesc = append(candlesDesc, klineToCandle(kl[i]))
	}

	yesterday := 1
	p := detectPattern(candlesDesc, yesterday, pair)
	return p
}
