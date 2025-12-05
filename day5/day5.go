package main

import (
	"day5/ingredients"
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
	ingredientRanges, ingredientIDs := ingredients.ParseIngredients(challengeText)
	totalFreshPartOne := 0
	for _, id := range ingredientIDs {
		for _, ingRange := range ingredientRanges {
			if ingRange.IsFresh(id) {
				totalFreshPartOne++
				break
			}
		}
	}

	fmt.Printf("Total fresh ingredients for Part 1: %d\n", totalFreshPartOne)
	// I could save time and just not do the above loop by creating some sort of
	// flag that would tell me not to count an ingredient again if it's already been counted,
	// but this is easier to grok i think.

	// This is wrong! I misunderstood the direction I should be going here. For
	// small ranges this is a no-brainer, but for larger ranges this is one of those O(log n)
	// problems that needs a sweeping line algorithm. So I implemented that.

	// Find all the fresh ingredient IDs
	// allIngs := ingredients.NewSet()
	// for _, ingRange := range ingredientRanges {
	// 	fmt.Printf("Debug: Processing range %s\n", ingRange.String())
	// 	ingRange.FindAllFreshIDs(allIngs)
	// }
	// fmt.Printf("Total fresh ingredients for Part 2: %d\n", allIngs.Size())

	totalFreshPartTwo := ingredients.SweepIngredientRanges(ingredientRanges)
	fmt.Printf("Total fresh ingredients for Part 2: %d\n", totalFreshPartTwo)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
