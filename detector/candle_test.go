package detector

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMorningStarOK(t *testing.T) {
	assert := assert.New(t)

	candles := _createDownTrend(200, 130, 7)
	candles = append(candles, Candle{Open: 130, Close: 120})           // red
	candles = append(candles, Candle{Open: 120, Close: 118, Low: 115}) // small
	candles = append(candles, Candle{Open: 118, Close: 125})           // green
	_reverse(candles)

	assert.Equal(true, IsMorningStar(candles))

	pattern := detectPattern(candles, 0, "anyway")
	assert.Equal("Morning Star", pattern.Type)
	assert.Equal("Bullish", pattern.TrendDirection)
}

func TestIsMorningStarDuringUpTrendShouldBeFalse(t *testing.T) {
	assert := assert.New(t)

	candles := _createUpTrend(30, 130, 10)
	candles = append(candles, Candle{Open: 130, Close: 120}) // red
	candles = append(candles, Candle{Open: 120, Close: 118}) // small
	candles = append(candles, Candle{Open: 118, Close: 125}) // green
	_reverse(candles)

	assert.Equal(false, IsMorningStar(candles))
}

func TestIsEveningStarOK(t *testing.T) {
	assert := assert.New(t)

	candles := _createUpTrend(50, 130, 7)
	candles = append(candles, Candle{Open: 130, Close: 140})                      // green
	candles = append(candles, Candle{Open: 140, Close: 138, Low: 135, High: 145}) // small with up wick
	candles = append(candles, Candle{Open: 138, Close: 120})                      // red
	_reverse(candles)

	assert.Equal(true, IsEveningStar(candles))

	pattern := detectPattern(candles, 0, "anyway")
	assert.Equal("Evening Star", pattern.Type)
	assert.Equal("Bearish", pattern.TrendDirection)
}

func TestIsEveningStarWithSmallWicksShouldBeFalse(t *testing.T) {
	assert := assert.New(t)

	candles := _createUpTrend(50, 130, 7)
	candles = append(candles, Candle{Open: 130, Close: 140})                      // green
	candles = append(candles, Candle{Open: 140, Close: 138, Low: 135, High: 142}) // small with small up wick
	candles = append(candles, Candle{Open: 138, Close: 120})                      // red
	_reverse(candles)

	assert.Equal(false, IsEveningStar(candles))
}

func TestIsEveningStarDuringDownTrendShouldBeFalse(t *testing.T) {
	assert := assert.New(t)

	candles := _createDownTrend(200, 130, 7)
	candles = append(candles, Candle{Open: 130, Close: 140})                      // green
	candles = append(candles, Candle{Open: 140, Close: 138, Low: 135, High: 145}) // small with up wick
	candles = append(candles, Candle{Open: 138, Close: 120})                      // red
	_reverse(candles)

	assert.Equal(false, IsEveningStar(candles))
}

func TestIsDojiWithNoWicksShouldBeFalse(t *testing.T) {
	assert := assert.New(t)
	var candles []Candle
	candles = append(candles, Candle{Open: 100, Close: 100.3}) // tiny candle with no wicks

	assert.Equal(false, IsDoji(candles))
}

func TestIsHammer(t *testing.T) {
	assert := assert.New(t)

	candles := _createDownTrend(220, 150, 10)
	candles = append(candles, Candle{Open: 150, Close: 145, Low: 122, High: 151.3})
	_reverse(candles)

	assert.Equal(true, IsHammer(candles))

	pattern := detectPattern(candles, 0, "anyway")
	assert.Equal("Hammer", pattern.Type)
	assert.Equal("Bullish", pattern.TrendDirection)
}

func TestIsHammerwithTooSmallBottomShadowShouldBeFalse(t *testing.T) {
	assert := assert.New(t)

	candles := _createDownTrend(220, 150, 10)
	candles = append(candles, Candle{Open: 150, Close: 145, Low: 140, High: 151.3})
	_reverse(candles)

	assert.Equal(false, IsHammer(candles))
}

func TestIsInvertedHammer(t *testing.T) {
	assert := assert.New(t)

	candles := _createUpTrend(100, 200, 10)
	candles = append(candles, Candle{Open: 200, Close: 195, Low: 194, High: 220})
	_reverse(candles)

	assert.Equal(true, IsInvertedHammer(candles))

	pattern := detectPattern(candles, 0, "anyway")
	assert.Equal("Inverted Hammer", pattern.Type)
	assert.Equal("Bearish", pattern.TrendDirection)
}

func TestIsWhiteMarubozu(t *testing.T) {
	assert := assert.New(t)

	candles := _createDownTrend(220, 150, 10)
	candles = append(candles, Candle{Open: 150, Close: 162, Low: 149, High: 163})
	_reverse(candles)

	assert.Equal(true, IsWhiteMarubozu(candles))

	pattern := detectPattern(candles, 0, "anyway")
	assert.Equal("White Marubozu", pattern.Type)
	assert.Equal("Bullish", pattern.TrendDirection)
}

func TestIsWhiteMarubozuWithShadowsShouldBeFalse(t *testing.T) {
	assert := assert.New(t)

	candles := _createDownTrend(220, 150, 10)
	candles = append(candles, Candle{Open: 150, Close: 162, Low: 146, High: 163})
	_reverse(candles)

	assert.Equal(false, IsWhiteMarubozu(candles))
}

func TestIsWhiteMarubozuDuringRangeShouldBeFalse(t *testing.T) {
	assert := assert.New(t)

	candles := _createUpTrend(140, 150, 10)
	candles = append(candles, Candle{Open: 150, Close: 162, Low: 149, High: 163})
	_reverse(candles)

	assert.Equal(false, IsWhiteMarubozu(candles))
}

func TestIsBlackMarubozu(t *testing.T) {
	assert := assert.New(t)

	candles := _createUpTrend(50, 150, 10)
	candles = append(candles, Candle{Open: 150, Close: 138, Low: 137, High: 151})
	_reverse(candles)

	assert.Equal(true, IsBlackMarubozu(candles))

	pattern := detectPattern(candles, 0, "anyway")
	assert.Equal("Black Marubozu", pattern.Type)
	assert.Equal("Bearish", pattern.TrendDirection)
}

func TestIsBlackMarubozuWithTopShadowShouldBeFalse(t *testing.T) {
	assert := assert.New(t)

	candles := _createUpTrend(50, 150, 10)
	candles = append(candles, Candle{Open: 150, Close: 138, Low: 138, High: 155})
	_reverse(candles)

	assert.Equal(false, IsBlackMarubozu(candles))
}

func TestIsDojiOK(t *testing.T) {
	assert := assert.New(t)
	var candles []Candle
	candles = append(candles, Candle{Open: 100, Close: 100.3, High: 108, Low: 92}) // tiny candle with wicks

	assert.Equal(true, IsDoji(candles))
}

func _reverse(candles []Candle) {
	for i, j := 0, len(candles)-1; i < j; i, j = i+1, j-1 {
		candles[i], candles[j] = candles[j], candles[i]
	}
}

func _createDownTrend(start, end float64, nbCandles int) []Candle {
	var candles []Candle
	step := math.Abs(end-start) / float64(nbCandles)
	for i := 0; i < nbCandles; i++ {
		candles = append(candles, Candle{Open: start, Close: start - step})
		start = start - step
	}
	return candles
}

func _createUpTrend(start, end float64, nbCandles int) []Candle {
	var candles []Candle
	step := math.Abs(end-start) / float64(nbCandles)
	for i := 0; i < nbCandles; i++ {
		candles = append(candles, Candle{Open: start, Close: start + step})
		start = start + step
	}
	return candles
}
