package main

import (
	"context"
	"fmt"
	"github.com/binance-exchange/go-binance"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
)

const baseUrl string = "https://api.binance.com"

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = log.With(logger, "time", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	//hmacSigner := &binance.HmacSigner{
	//	Key: []byte(os.Getenv("BINANCE_SECRET")),
	//}
	ctx, _ := context.WithCancel(context.Background())
	// use second return value for cancelling request

	binanceService := binance.NewAPIService(
		"https://www.binance.com",
		"",  //os.Getenv("BINANCE_APIKEY"),
		nil, //hmacSigner,
		logger,
		ctx,
	)
	b := binance.NewBinance(binanceService)

	kl, err := b.Klines(binance.KlinesRequest{
		Symbol:   "BTCUSDT",
		Interval: binance.Hour,
		Limit:    50,
	})
	if err != nil {
		panic(err)
	}

	for _, kline := range kl {
		c := Candle{kline}
		fmt.Printf("%#v - %#.2f%% \n", kline.Open, c.BodyPercent())
	}

}
