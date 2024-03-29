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

func (candle Candle) isRed() bool   { return candle.Open > candle.Close }
func (candle Candle) isGreen() bool { return candle.Open < candle.Close }

//body size of the body
func (candle Candle) body() float64 { return math.Abs(candle.Open - candle.Close) }

//bodyRatio ratio with open price
func (candle Candle) bodyRatio() float64 { return math.Abs(candle.Open-candle.Close) / candle.Open }

//bodyPercent body ratio in percent
func (candle Candle) bodyPercent() float64 { return candle.bodyRatio() * 100 }

//bodyLowest the lowest price in body's candle
func (candle Candle) bodyLowest() float64 { return math.Min(candle.Open, candle.Close) }

//bodyHighest the highest price in body's candle
func (candle Candle) bodyHighest() float64 { return math.Max(candle.Open, candle.Close) }

func (candle Candle) topShadowRatio() float64 {
	if candle.High == 0 {
		candle.High = candle.bodyHighest()
	}

	if candle.isRed() {
		return math.Abs(candle.High-candle.Open) / candle.Open
	} else {
		return math.Abs(candle.High-candle.Close) / candle.Close
	}
}

func (candle Candle) bottomShadowRatio() float64 {
	if candle.Low == 0 {
		candle.Low = candle.bodyLowest()
	}

	if candle.isRed() {
		return math.Abs(candle.Close-candle.Low) / candle.Close
	} else {
		return math.Abs(candle.Open-candle.Low) / candle.Open
	}
}

//func (candle RedCandle) LowShadow() float64 { return candle.Low - candle.Close }
//func (candle GreenCandle) LowShadow() float64 { return candle.Low - candle.Open }

func (candle Candle) isTiny() bool  { return candle.bodyRatio() < maxTinyBody }
func (candle Candle) isSmall() bool { return candle.bodyRatio() < maxSmallBody }
func (candle Candle) isBody() bool {
	return maxSmallBody < candle.bodyRatio() && candle.bodyRatio() < minBigBody
}
func (candle Candle) isLong() bool { return candle.bodyRatio() > minBigBody }

func (candle Candle) hasTinyTopShadow() bool  { return candle.topShadowRatio() < maxTinyBody }
func (candle Candle) hasTopShadow() bool      { return candle.topShadowRatio() > maxSmallBody }
func (candle Candle) hasSmallTopShadow() bool { return true }
func (candle Candle) hasLongTopShadow() bool  { return candle.topShadowRatio() > minBigBody }

func (candle Candle) hasTinyBottomShadow() bool  { return candle.bottomShadowRatio() < maxTinyBody }
func (candle Candle) hasBottomShadow() bool      { return candle.bottomShadowRatio() > maxSmallBody }
func (candle Candle) hasSmallBottomShadow() bool { return true }
func (candle Candle) hasLongBottomShadow() bool  { return candle.bottomShadowRatio() > minBigBody }

func IsMorningStar(candlesSortDesc []Candle) bool {
	// one red not small plus one small with long low shadow and green not small and during a downtrend
	return len(candlesSortDesc) > 2 && candlesSortDesc[0].isGreen() && !candlesSortDesc[0].isSmall() &&
		candlesSortDesc[1].isSmall() && candlesSortDesc[1].hasBottomShadow() &&
		candlesSortDesc[2].isRed() && !candlesSortDesc[2].isSmall() &&
		IsDownTrend(candlesSortDesc[2:], 5)
}

func IsEveningStar(candlesSortDesc []Candle) bool {
	// one green not small plus one small with long low shadow and a red not small and during an uptrend
	return len(candlesSortDesc) > 2 &&
		candlesSortDesc[0].isRed() && !candlesSortDesc[0].isSmall() &&
		candlesSortDesc[1].isSmall() && candlesSortDesc[1].hasTopShadow() &&
		candlesSortDesc[2].isGreen() && !candlesSortDesc[2].isSmall() &&
		IsUpTrend(candlesSortDesc[2:], 5)
}

func IsThreeWhiteSoldiers(candlesSortDesc []Candle) bool {
	// three green with no or very small shadow and during a downtrend
	return len(candlesSortDesc) > 2 &&
		candlesSortDesc[0].isGreen() && !candlesSortDesc[0].isSmall() &&
		candlesSortDesc[1].isGreen() && !candlesSortDesc[1].isSmall() &&
		candlesSortDesc[2].isGreen() && !candlesSortDesc[2].isSmall() &&
		IsDownTrend(candlesSortDesc[2:], 9)
}

func IsThreeBlackCrows(candlesSortDesc []Candle) bool {
	// three red with no or very small shadow and during an uptrend
	return len(candlesSortDesc) > 2 &&
		candlesSortDesc[0].isRed() && !candlesSortDesc[0].isSmall() &&
		candlesSortDesc[1].isRed() && !candlesSortDesc[1].isSmall() &&
		candlesSortDesc[2].isRed() && !candlesSortDesc[2].isSmall() &&
		IsUpTrend(candlesSortDesc[2:], 9)
}

func IsWhiteMarubozu(candlesSortDesc []Candle) bool {
	// a long green with no or very small shadow and during an downtrend
	return len(candlesSortDesc) > 1 &&
		candlesSortDesc[0].isGreen() && candlesSortDesc[0].isLong() &&
		candlesSortDesc[0].hasTinyBottomShadow() && candlesSortDesc[0].hasTinyTopShadow() &&
		IsDownTrend(candlesSortDesc[2:], 7)
}

func IsBlackMarubozu(candlesSortDesc []Candle) bool {
	// a long red with no or very small shadow and during an uptrend
	return len(candlesSortDesc) > 1 &&
		candlesSortDesc[0].isRed() && candlesSortDesc[0].isLong() &&
		candlesSortDesc[0].hasTinyBottomShadow() && candlesSortDesc[0].hasTinyTopShadow() &&
		IsUpTrend(candlesSortDesc[2:], 7)
}

func IsDoji(candleSortDesc []Candle) bool {
	// a very small (tiny) with long top and bottom shadows
	return len(candleSortDesc) > 0 &&
		candleSortDesc[0].isTiny() &&
		candleSortDesc[0].hasBottomShadow() && candleSortDesc[0].hasTopShadow()
}

func IsHammer(candleSortDesc []Candle) bool {
	// a real body with no or little top shadow and long (twice body) bottom shadow, during a downtrend
	return len(candleSortDesc) > 0 &&
		candleSortDesc[0].isBody() &&
		candleSortDesc[0].hasLongBottomShadow() && candleSortDesc[0].hasTinyTopShadow() &&
		IsDownTrend(candleSortDesc[1:], 7)
}

func IsInvertedHammer(candleSortDesc []Candle) bool {
	// a real body with no or little bottom shadow and long (twice body) top shadow, during an uptrend
	return len(candleSortDesc) > 0 &&
		candleSortDesc[0].isBody() &&
		candleSortDesc[0].hasLongTopShadow() && candleSortDesc[0].hasTinyBottomShadow() &&
		IsUpTrend(candleSortDesc[1:], 7)
}

// IsDownTrend is in a current downtrend
// remarks:  not good, but it's a first shot
func IsDownTrend(candlesSortDesc []Candle, trendDuration int) bool {
	return len(candlesSortDesc) > trendDuration &&
		candlesSortDesc[0].bodyLowest() < candlesSortDesc[trendDuration].bodyLowest()
}

// IsUpTrend is in a current uptrend
// remarks:  not good, but it's a first shot
func IsUpTrend(candlesSortDesc []Candle, trendDuration int) bool {
	return len(candlesSortDesc) > trendDuration &&
		candlesSortDesc[0].bodyLowest() > candlesSortDesc[trendDuration].bodyLowest()
}
