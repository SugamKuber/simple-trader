package analysis

import (
	"fmt"
	"trader/types"
	"trader/utils"
)

func FindTradingSignals(data []types.OHCLV, minWeight float64) []string {
	var signals []string

	for i := 3; i < len(data); i++ {
		candles := data[i-3 : i+1]

		allGreen := true
		for _, candle := range candles {
			if !IsGreenCandle(candle) {
				allGreen = false
				break
			}
		}

		hasStairPattern, weight := HasValidStairStepPattern(candles)

		if allGreen && hasStairPattern && weight >= minWeight {
			signal := fmt.Sprintf("Signal found at: %s UTC\n"+
				"Current Price: %.8f\n"+
				"Pattern Weight: %.4f\n",
				utils.EpochMillisToTime(candles[3].Time),
				candles[3].Close,
				weight)

			signals = append(signals, signal)
		}
	}

	return signals
}
