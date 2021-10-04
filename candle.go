package main

import "github.com/binance-exchange/go-binance"

type Candle struct {
	*binance.Kline
}

func (candle Candle) IsRed() bool {
	return candle.Close < candle.Open
}

func (candle Candle) IsGreen() bool {
	return candle.Close > candle.Open
}

func (candle Candle) IsSmall() bool {
	return false
}

func (candle Candle) IsBody() bool {
	return false
}

func (candle Candle) IsLong() bool {
	return false
}
