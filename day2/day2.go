package main

import (
	"day2/ranges"
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
	// lines, err := utils.GetFileLines(*filePtr)
	// if err != nil {
	//     fmt.Println("Fatal:", err)
	// }

	// giant text blob:
	challengeText, err := utils.GetTextBlob(*filePtr)
	if err != nil {
		fmt.Println("Fatal:", err)
	}

	// Insert code here
	rangeBlob := ranges.ParseRanges(challengeText)
	totalRanges := []ranges.Range{}
	for _, rangeText := range rangeBlob {
		r, err := ranges.NewRange(rangeText)
		if err != nil {
			fmt.Println("Error parsing range:", err)
			continue
		}
		totalRanges = append(totalRanges, *r)
	}

	totalInvalid := 0
	totalInvalidPart2 := 0
	for _, r := range totalRanges {
		invalidRanges := r.GetAllInvalidIds()
		for _, inv := range invalidRanges {
			// fmt.Printf("Invalid ID found: %d\n", inv)
			totalInvalid += inv
		}

		invalidRangesPart2 := r.GetAllInvalidIdsPart2()
		for _, inv := range invalidRangesPart2 {
			totalInvalidPart2 += inv
		}
	}
	fmt.Printf("Invalid total in all ranges: %d\n", totalInvalid)
	fmt.Printf("Invalid total in all ranges for part 2: %d\n", totalInvalidPart2)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
