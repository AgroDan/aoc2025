package main

import (
	"day3/batterybank"
	"flag"
	"fmt"
	"time"
	"utils"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")
	// any additional flags add here

	flag.Parse()

	// Choose based on the challenge...
	// individual lines:
	lines, err := utils.GetFileLines(*filePtr)
	if err != nil {
		fmt.Println("Fatal:", err)
	}

	// giant text blob:
	// challengeText, err := utils.GetTextBlob(*filePtr)
	// if err != nil {
	//     fmt.Println("Fatal:", err)
	// }

	// Insert code here

	batteries := make([]*batterybank.BatteryBank, 0)
	for _, line := range lines {
		batteries = append(batteries, batterybank.NewBatteryBank(line))
	}

	// For yuks lets just loop over it all again
	totalPartOne := 0
	totalPartTwo := 0
	for _, bank := range batteries {
		// fmt.Printf("Bank %d: Total: %d, Cells: %v\n", i, bank.TotalCharge(), bank.Cells)
		totalPartOne += bank.TotalCharge()
		totalPartTwo += bank.GetHighestPossible(12)
	}
	fmt.Printf("Total Charge from all banks Part 1: %d\n", totalPartOne)
	fmt.Printf("Total Charge from all banks Part 2: %d\n", totalPartTwo)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
