package detectdaily

import (
	"math"
	"time"
)

const maxTinyBody float64 = 0.01
const maxSmallBody float64 = 0.02
const minBigBody float64 = 0.05

//const maxTailShadow float64 = 0.05
//const minLongShadow float64 = 0.6

type Candle struct {
	OpenTime                 time.Time
	Open                     float64
	High                     float64
	Low                      float64
	Close                    float64
	Volume                   float64
	CloseTime                time.Time
	QuoteAssetVolume         float64
	NumberOfTrades           int
	TakerBuyBaseAssetVolume  float64
	TakerBuyQuoteAssetVolume float64
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

func (candle Candle) TopShadowRatio() float64 {
	if candle.IsRed() {
		return (candle.High - candle.Open) / candle.Open
	} else {
		return (candle.High - candle.Close) / candle.Close
	}
}

func (candle Candle) BottomShadowRatio() float64 {
	if candle.IsRed() {
		return (candle.Close - candle.Low) / candle.Close
	} else {
		return (candle.Open - candle.Low) / candle.Open
	}
}

//func (candle RedCandle) LowShadow() float64 { return candle.Low - candle.Close }
//func (candle GreenCandle) LowShadow() float64 { return candle.Low - candle.Open }

func (candle Candle) IsTiny() bool  { return candle.BodyRatio() < maxTinyBody }
func (candle Candle) IsSmall() bool { return candle.BodyRatio() < maxSmallBody }
func (candle Candle) IsBody() bool {
	return maxSmallBody > candle.BodyRatio() && candle.BodyRatio() < minBigBody
}
func (candle Candle) IsLong() bool { return candle.BodyRatio() > minBigBody }

func (candle Candle) HasTinyTopShadow() bool  { return true }
func (candle Candle) HasTopShadow() bool      { return candle.TopShadowRatio() > maxSmallBody }
func (candle Candle) HasSmallTopShadow() bool { return true }
func (candle Candle) HasLongTopShadow() bool  { return true }

func (candle Candle) HasTinyBottomShadow() bool  { return true }
func (candle Candle) HasBottomShadow() bool      { return candle.BottomShadowRatio() > maxSmallBody }
func (candle Candle) HasSmallBottomShadow() bool { return true }
func (candle Candle) HasLongBottomShadow() bool  { return true }

func IsMorningStar(candlesSortDesc []Candle) bool {
	// one red not small plus one small with long low shadow and green not small and during a downtrend
	return len(candlesSortDesc) > 2 && candlesSortDesc[0].IsGreen() && !candlesSortDesc[0].IsSmall() &&
		candlesSortDesc[1].IsSmall() && candlesSortDesc[1].HasBottomShadow() &&
		candlesSortDesc[2].IsRed() && !candlesSortDesc[2].IsSmall() &&
		IsDownTrend(candlesSortDesc[2:], 5)
}

func IsEveningStar(candlesSortDesc []Candle) bool {
	// one green not small plus one small with long low shadow and a red not small and during an uptrend
	return len(candlesSortDesc) > 2 &&
		candlesSortDesc[0].IsRed() && !candlesSortDesc[0].IsSmall() &&
		candlesSortDesc[1].IsSmall() && candlesSortDesc[1].HasTopShadow() &&
		candlesSortDesc[2].IsGreen() && !candlesSortDesc[2].IsSmall() &&
		IsUpTrend(candlesSortDesc[2:], 5)
}

func IsThreeWhiteSoldiers(candlesSortDesc []Candle) bool {
	// three green with no or very small shadow and during a downtrend
	return len(candlesSortDesc) > 2 &&
		candlesSortDesc[0].IsGreen() && !candlesSortDesc[0].IsSmall() &&
		candlesSortDesc[1].IsGreen() && !candlesSortDesc[1].IsSmall() &&
		candlesSortDesc[2].IsGreen() && !candlesSortDesc[2].IsSmall() &&
		IsDownTrend(candlesSortDesc[2:], 9)
}

func IsThreeBlackCrows(candlesSortDesc []Candle) bool {
	// three red with no or very small shadow and during an uptrend
	return len(candlesSortDesc) > 2 &&
		candlesSortDesc[0].IsRed() && !candlesSortDesc[0].IsSmall() &&
		candlesSortDesc[1].IsRed() && !candlesSortDesc[1].IsSmall() &&
		candlesSortDesc[2].IsRed() && !candlesSortDesc[2].IsSmall() &&
		IsUpTrend(candlesSortDesc[2:], 9)
}

func IsWhiteMarubozu(candlesSortDesc []Candle) bool {
	// a long green with no or very small shadow and during an downtrend
	return len(candlesSortDesc) > 1 &&
		candlesSortDesc[0].IsGreen() && candlesSortDesc[0].IsLong() &&
		IsDownTrend(candlesSortDesc[2:], 7)
}

func IsBlackMarubozu(candlesSortDesc []Candle) bool {
	// a long red with no or very small shadow and during an uptrend
	return len(candlesSortDesc) > 1 &&
		candlesSortDesc[0].IsRed() && candlesSortDesc[0].IsLong() &&
		IsUpTrend(candlesSortDesc[2:], 7)
}

func IsDoji(candleSortDesc []Candle) bool {
	// a very small (tiny) with long top and bottom shadows
	return len(candleSortDesc) > 0 &&
		candleSortDesc[0].IsTiny() &&
		candleSortDesc[0].HasBottomShadow() && candleSortDesc[0].HasTopShadow()
}

// IsDownTrend is in a current downtrend
// remarks:  not good, but it's a first shot
func IsDownTrend(candlesSortDesc []Candle, trendDuration int) bool {
	return len(candlesSortDesc) > trendDuration &&
		candlesSortDesc[0].BodyLowest() < candlesSortDesc[trendDuration].BodyLowest()
}

// IsUpTrend is in a current uptrend
// remarks:  not good, but it's a first shot
func IsUpTrend(candlesSortDesc []Candle, trendDuration int) bool {
	return len(candlesSortDesc) > trendDuration &&
		candlesSortDesc[0].BodyLowest() > candlesSortDesc[trendDuration].BodyLowest()
}
