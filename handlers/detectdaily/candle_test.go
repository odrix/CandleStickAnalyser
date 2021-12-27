package detectdaily

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMorningStarOK(t *testing.T) {
	assert := assert.New(t)

	candles := createDownTrend(200, 130, 7)
	candles = append(candles, Candle{Open: 130, Close: 120}) // red
	candles = append(candles, Candle{Open: 120, Close: 118}) // small
	candles = append(candles, Candle{Open: 118, Close: 125}) // green
	Reverse(candles)

	assert.Equal(true, IsMorningStar(candles))
}

func TestIsMorningStarDuringUpTrendShouldBeFalse(t *testing.T) {
	assert := assert.New(t)

	candles := createUpTrend(30, 130, 10)
	candles = append(candles, Candle{Open: 130, Close: 120}) // red
	candles = append(candles, Candle{Open: 120, Close: 118}) // small
	candles = append(candles, Candle{Open: 118, Close: 125}) // green
	Reverse(candles)

	assert.Equal(false, IsMorningStar(candles))
}

func Reverse(candles []Candle) {
	for i, j := 0, len(candles)-1; i < j; i, j = i+1, j-1 {
		candles[i], candles[j] = candles[j], candles[i]
	}
}

func createDownTrend(start, end float64, nbCandles int) []Candle {
	var candles []Candle
	step := math.Abs(end-start) / float64(nbCandles)
	for i := 0; i < nbCandles; i++ {
		candles = append(candles, Candle{Open: start, Close: start - step})
		start = start - step
	}
	return candles
}

func createUpTrend(start, end float64, nbCandles int) []Candle {
	var candles []Candle
	step := math.Abs(end-start) / float64(nbCandles)
	for i := 0; i < nbCandles; i++ {
		candles = append(candles, Candle{Open: start, Close: start + step})
		start = start + step
	}
	return candles
}
