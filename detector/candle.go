package detector

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
	if candle.High == 0 {
		candle.High = candle.BodyHighest()
	}

	if candle.IsRed() {
		return math.Abs(candle.High-candle.Open) / candle.Open
	} else {
		return math.Abs(candle.High-candle.Close) / candle.Close
	}
}

func (candle Candle) BottomShadowRatio() float64 {
	if candle.Low == 0 {
		candle.Low = candle.BodyLowest()
	}

	if candle.IsRed() {
		return math.Abs(candle.Close-candle.Low) / candle.Close
	} else {
		return math.Abs(candle.Open-candle.Low) / candle.Open
	}
}

//func (candle RedCandle) LowShadow() float64 { return candle.Low - candle.Close }
//func (candle GreenCandle) LowShadow() float64 { return candle.Low - candle.Open }

func (candle Candle) IsTiny() bool  { return candle.BodyRatio() < maxTinyBody }
func (candle Candle) IsSmall() bool { return candle.BodyRatio() < maxSmallBody }
func (candle Candle) IsBody() bool {
	return maxSmallBody < candle.BodyRatio() && candle.BodyRatio() < minBigBody
}
func (candle Candle) IsLong() bool { return candle.BodyRatio() > minBigBody }

func (candle Candle) HasTinyTopShadow() bool  { return candle.TopShadowRatio() < maxTinyBody }
func (candle Candle) HasTopShadow() bool      { return candle.TopShadowRatio() > maxSmallBody }
func (candle Candle) HasSmallTopShadow() bool { return true }
func (candle Candle) HasLongTopShadow() bool  { return candle.TopShadowRatio() > minBigBody }

func (candle Candle) HasTinyBottomShadow() bool  { return candle.BottomShadowRatio() < maxTinyBody }
func (candle Candle) HasBottomShadow() bool      { return candle.BottomShadowRatio() > maxSmallBody }
func (candle Candle) HasSmallBottomShadow() bool { return true }
func (candle Candle) HasLongBottomShadow() bool  { return candle.BottomShadowRatio() > minBigBody }

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
		candlesSortDesc[0].HasTinyBottomShadow() && candlesSortDesc[0].HasTinyTopShadow() &&
		IsDownTrend(candlesSortDesc[2:], 7)
}

func IsBlackMarubozu(candlesSortDesc []Candle) bool {
	// a long red with no or very small shadow and during an uptrend
	return len(candlesSortDesc) > 1 &&
		candlesSortDesc[0].IsRed() && candlesSortDesc[0].IsLong() &&
		candlesSortDesc[0].HasTinyBottomShadow() && candlesSortDesc[0].HasTinyTopShadow() &&
		IsUpTrend(candlesSortDesc[2:], 7)
}

func IsDoji(candleSortDesc []Candle) bool {
	// a very small (tiny) with long top and bottom shadows
	return len(candleSortDesc) > 0 &&
		candleSortDesc[0].IsTiny() &&
		candleSortDesc[0].HasBottomShadow() && candleSortDesc[0].HasTopShadow()
}

func IsHammer(candleSortDesc []Candle) bool {
	// a real body with no or little top shadow and long (twice body) bottom shadow, during a downtrend
	return len(candleSortDesc) > 0 &&
		candleSortDesc[0].IsBody() &&
		candleSortDesc[0].HasLongBottomShadow() && candleSortDesc[0].HasTinyTopShadow() &&
		IsDownTrend(candleSortDesc[1:], 7)
}

func IsInvertedHammer(candleSortDesc []Candle) bool {
	// a real body with no or little bottom shadow and long (twice body) top shadow, during an uptrend
	return len(candleSortDesc) > 0 &&
		candleSortDesc[0].IsBody() &&
		candleSortDesc[0].HasLongTopShadow() && candleSortDesc[0].HasTinyBottomShadow() &&
		IsUpTrend(candleSortDesc[1:], 7)
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
