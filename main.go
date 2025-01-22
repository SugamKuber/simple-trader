package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type OHCLV struct {
	Time   int64
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

func isGreenCandle(candle OHCLV) bool {
	return candle.Close > candle.Open
}

func hasValidStairStepPattern(candles []OHCLV) bool {
	if len(candles) < 2 {
		return false
	}

	for i := 1; i < len(candles); i++ {
		current := candles[i]
		previous := candles[i-1]

		openGap := current.Open - previous.Close
		closeHigher := current.Close > previous.Close

		maxAllowedGap := -0.2 * math.Abs(previous.Close-previous.Open)

		if openGap < maxAllowedGap || !closeHigher {
			return false
		}
	}
	return true
}

func findTradingSignals(data []OHCLV) []string {
	var signals []string

	for i := 3; i < len(data); i++ {

		candles := data[i-3 : i+1]

		allGreen := true
		for _, candle := range candles {
			if !isGreenCandle(candle) {
				allGreen = false
				break
			}
		}

		hasStairPattern := hasValidStairStepPattern(candles)

		if allGreen && hasStairPattern {
			signal := fmt.Sprintf("Signal found at %s:\n"+
				"Current Price: %.8f\n",
				epochMillisToTime(candles[3].Time),
				candles[3].Close)

			signals = append(signals, signal)
		}
	}

	return signals
}

func epochMillisToTime(epochMillis int64) string {
	t := time.Unix(0, epochMillis*int64(time.Millisecond))
	return t.Format("2006-01-02 15:04:05.000 MST")
}

func main() {
	file, err := os.Open("BTCUSDT-1h-2024-12.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	var data []OHCLV
	for _, record := range records {
		time, _ := strconv.ParseInt(record[0], 10, 64)
		open, _ := strconv.ParseFloat(record[1], 64)
		high, _ := strconv.ParseFloat(record[2], 64)
		low, _ := strconv.ParseFloat(record[3], 64)
		close, _ := strconv.ParseFloat(record[4], 64)
		volume, _ := strconv.ParseFloat(record[5], 64)
		data = append(data, OHCLV{Time: time, Open: open, High: high, Low: low, Close: close, Volume: volume})
	}

	signals := findTradingSignals(data)

	for _, signal := range signals {
		fmt.Print(signal)
		fmt.Println("-------------------")
	}
}
