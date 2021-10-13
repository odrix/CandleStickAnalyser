package main

import (
	"github.com/binance-exchange/go-binance"
	"math"
)

const maxCloseBody float64 = 0.005
const maxSmallBody float64 = 0.02
const minBigBody float64 = 0.05

//const maxTailShadow float64 = 0.05
//const minLongShadow float64 = 0.6

type Candle struct {
	*binance.Kline
}

func (candle Candle) IsRed() bool   { return candle.Open > candle.Close }
func (candle Candle) IsGreen() bool { return candle.Open < candle.Close }

//Body size of the body
func (candle Candle) Body() float64 { return math.Abs(candle.Open - candle.Close) }

//BodyRatio ratio with open price
func (candle Candle) BodyRatio() float64 { return math.Abs(candle.Open-candle.Close) / candle.Open }

//BodyPercent body ratio in percent
func (candle Candle) BodyPercent() float64 { return candle.BodyRatio() * 100 }

//BodyLowest the lowest price in body's candle
func (candle Candle) BodyLowest() float64 { return math.Min(candle.Open, candle.Close) }

//BodyHighest the highest price in body's candle
func (candle Candle) BodyHighest() float64 { return math.Max(candle.Open, candle.Close) }

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

func IsMorningStar(candlesSortDesc []Candle) bool {
	// one red not small plus one small with long low shadow and green not small and during a downtrend
	return candlesSortDesc[0].IsGreen() && !candlesSortDesc[0].IsSmall() &&
		candlesSortDesc[1].IsSmall() &&
		candlesSortDesc[2].IsRed() && !candlesSortDesc[2].IsSmall() &&
		IsDownTrend(candlesSortDesc[2:])

}

// IsDownTrend is in a current downtrend
// remarks:  not good, but it's a first shot
func IsDownTrend(candlesSortDesc []Candle) bool {
	return candlesSortDesc[0].BodyLowest() < candlesSortDesc[5].BodyLowest() ||
		candlesSortDesc[0].BodyLowest() < candlesSortDesc[8].BodyLowest()
}
