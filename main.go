package main

import (
	"context"
	"fmt"
	"github.com/binance-exchange/go-binance"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
)

//const baseUrl string = "https://api.binance.com"

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = log.With(logger, "time", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

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
		Symbol:   "BTCUSDT",
		Interval: binance.Day,
		Limit:    300,
	})
	if err != nil {
		panic(err)
	}

	var candlesDesc []Candle
	for i := len(kl) - 1; i > 0; i-- {
		candlesDesc = append(candlesDesc, Candle{kl[i]})
	}

	for i := 0; i < len(candlesDesc); i++ {
		if IsMorningStar(candlesDesc[i:]) {
			fmt.Printf("morning star on %s \n", candlesDesc[i].OpenTime)
		}
	}
	fmt.Printf("end.")

}
