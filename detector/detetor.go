package detector

import (
	"context"
	"fmt"

	"github.com/binance-exchange/go-binance"
	"github.com/go-kit/kit/log"
)

func DetectPattern(candlesDesc []Candle, startDayIndex int, pair string) Pattern {
	p := Pattern{}
	if IsMorningStar(candlesDesc[startDayIndex:]) {
		p = Pattern{
			Pair:           pair,
			Type:           "Morning Star",
			TrendDirection: "Bullish",
			Start:          candlesDesc[startDayIndex+2].OpenTime,
			End:            candlesDesc[startDayIndex].CloseTime,
		}
	} else if IsEveningStar(candlesDesc[startDayIndex:]) {
		p = Pattern{
			Pair:           pair,
			Type:           "Evening Star",
			TrendDirection: "Bearish",
			Start:          candlesDesc[startDayIndex+2].OpenTime,
			End:            candlesDesc[startDayIndex].CloseTime,
		}
	} else if IsThreeWhiteSoldiers(candlesDesc[startDayIndex:]) {
		p = Pattern{
			Pair:           pair,
			Type:           "Three White Soldiers",
			TrendDirection: "Bullish",
			Start:          candlesDesc[startDayIndex+2].OpenTime,
			End:            candlesDesc[startDayIndex].CloseTime,
		}
	} else if IsThreeBlackCrows(candlesDesc[startDayIndex:]) {
		p = Pattern{
			Pair:           pair,
			Type:           "Three Black Crows",
			TrendDirection: "Bearish",
			Start:          candlesDesc[startDayIndex+2].OpenTime,
			End:            candlesDesc[startDayIndex].CloseTime,
		}
	} else if IsWhiteMarubozu(candlesDesc[startDayIndex:]) {
		p = Pattern{
			Pair:           pair,
			Type:           "White Marubozu",
			TrendDirection: "Bullish",
			Start:          candlesDesc[startDayIndex].OpenTime,
			End:            candlesDesc[startDayIndex].CloseTime,
		}
	} else if IsBlackMarubozu(candlesDesc[startDayIndex:]) {
		p = Pattern{
			Pair:           pair,
			Type:           "Black Marubozu",
			TrendDirection: "Bearish",
			Start:          candlesDesc[startDayIndex].OpenTime,
			End:            candlesDesc[startDayIndex].CloseTime,
		}
	} else if IsHammer(candlesDesc[startDayIndex:]) {
		p = Pattern{
			Pair:           pair,
			Type:           "Hammer",
			TrendDirection: "Bullish",
			Start:          candlesDesc[startDayIndex].OpenTime,
			End:            candlesDesc[startDayIndex].CloseTime,
		}
	} else if IsInvertedHammer(candlesDesc[startDayIndex:]) {
		p = Pattern{
			Pair:           pair,
			Type:           "Inverted Hammer",
			TrendDirection: "Bearish",
			Start:          candlesDesc[startDayIndex].OpenTime,
			End:            candlesDesc[startDayIndex].CloseTime,
		}
	} else if IsDoji(candlesDesc[startDayIndex:]) {
		p = Pattern{
			Pair:           pair,
			Type:           "Doji",
			TrendDirection: "Continuation",
			Start:          candlesDesc[startDayIndex].OpenTime,
			End:            candlesDesc[startDayIndex].CloseTime,
		}
	}
	return p
}

func GetKline(logger log.Logger, symbol string, interval binance.Interval, limit int) []*binance.Kline {
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

func KlineToCandle(k *binance.Kline) Candle {
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
