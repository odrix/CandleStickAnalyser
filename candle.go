package main

import (
	"github.com/binance-exchange/go-binance"
	"math"
)

const maxCloseBody float64 = 0.005
const maxSmallBody float64 = 0.02
const minBigBody float64 = 0.05
const maxTailShadow float64 = 0.05
const minLongShadow float64 = 0.6

//type Candle interface{
//	HighShadow() float64
//	LowShadow() float64
//	Body() float64
//	type RedCandle, GreenCandle
//}

//type RedCandle struct {
//	*binance.Kline
//}
//
//type GreenCandle struct {
//	*binance.Kline
//}

type Candle struct {
	*binance.Kline
}

//func CreateCandle(kline *binance.Kline) Candle {
//	if kline.Close < kline.Open {
//		return  RedCandle{kline}
//	} else {
//		return  GreenCandle{kline}
//	}
//}

func (candle Candle) Body() float64        { return math.Abs(candle.Open - candle.Close) }
func (candle Candle) BodyRatio() float64   { return math.Abs(candle.Open-candle.Close) / candle.Open }
func (candle Candle) BodyPercent() float64 { return candle.BodyRatio() * 100 }

//func (candle RedCandle) LowShadow() float64 { return candle.Low - candle.Close }
//func (candle GreenCandle) LowShadow() float64 { return candle.Low - candle.Open }
//
//func (candle RedCandle) HighShadow() float64 { return candle.High - candle.Open }
//func (candle GreenCandle) HighShadow() float64 { return candle.High - candle.Close }

func (candle Candle) IsClose() bool { return candle.BodyRatio() < maxCloseBody }
func (candle Candle) IsSmall() bool { return candle.BodyRatio() < maxSmallBody }
func (candle Candle) IsBody() bool {
	return maxSmallBody > candle.BodyRatio() && candle.BodyRatio() < minBigBody
}
func (candle Candle) IsLong() bool { return candle.BodyRatio() > minBigBody }

func IsMorningStar(candles []Candle) bool {
	return false // one red not small plus one small with long low shadow and green not small
}
