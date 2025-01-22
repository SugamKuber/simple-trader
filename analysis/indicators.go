package analysis

import (
	"math"
	"trader/types"
)

func GetTrueRange(current, previous types.OHCLV) float64 {
	highLow := math.Abs(current.High - current.Low)
	highClose := math.Abs(current.High - previous.Close)
	lowClose := math.Abs(current.Low - previous.Close)
	return math.Max(highLow, math.Max(highClose, lowClose))
}

func GetATR(candles []types.OHCLV, period int) float64 {
	if len(candles) < period+1 {
		return 0
	}

	var sum float64
	for i := 1; i <= period; i++ {
		tr := GetTrueRange(candles[i], candles[i-1])
		sum += tr
	}
	return sum / float64(period)
}
