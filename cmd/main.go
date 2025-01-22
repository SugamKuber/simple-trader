package main

import (
	"fmt"
	"trader/analysis"
	"trader/data"
)

func main() {
	data, err := data.ReadCSVData("BTCUSDT-1h-2024-12.csv")
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	minWeight := 0.5
	signals := analysis.FindTradingSignals(data, minWeight)

	for _, signal := range signals {
		fmt.Print(signal)
		fmt.Println("-------------------")
	}
}
