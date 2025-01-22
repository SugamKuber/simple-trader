package analysis

import (
	"math"
	"trader/types"
)

func IsGreenCandle(candle types.OHCLV) bool {
	return candle.Close > candle.Open
}

func GetPatternWeight(candles []types.OHCLV) float64 {
	weight := 1.0

	atr := GetATR(candles, len(candles)-1)
	avgPrice := candles[len(candles)-1].Close
	volatility := atr / avgPrice

	volatilityWeight := 1.0 - math.Min(volatility*10, 0.8)
	weight *= volatilityWeight

	var priceConsistency float64
	for i := 1; i < len(candles); i++ {
		current := candles[i]
		previous := candles[i-1]

		expectedMove := (current.Close - previous.Close) / previous.Close
		idealMove := 0.01
		consistency := 1.0 - math.Min(math.Abs(expectedMove-idealMove)/idealMove, 0.5)
		priceConsistency += consistency
	}
	priceConsistency /= float64(len(candles) - 1)
	weight *= priceConsistency

	return weight
}

func HasValidStairStepPattern(candles []types.OHCLV) (bool, float64) {
	if len(candles) < 2 {
		return false, 0
	}

	patternValid := true
	for i := 1; i < len(candles); i++ {
		current := candles[i]
		previous := candles[i-1]

		openGap := current.Open - previous.Close
		closeHigher := current.Close > previous.Close

		maxAllowedGap := -0.2 * math.Abs(previous.Close-previous.Open)

		if openGap < maxAllowedGap || !closeHigher {
			patternValid = false
			break
		}
	}

	weight := 0.0
	if patternValid {
		weight = GetPatternWeight(candles)
	}

	return patternValid, weight
}
