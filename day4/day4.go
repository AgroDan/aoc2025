package main

import (
	"day4/rolls"
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

	// This loads everything in as a runemap, so I'll just use that
	papermap := utils.NewRunemap(lines)
	papermap.Print()

	accessibleRolls := rolls.FindAccessibleRolls(papermap)
	fmt.Println("Accessible Rolls for part 1:", len(accessibleRolls))

	removed := rolls.RemoveRolls(&papermap)

	papermap.Print()
	fmt.Println("Total rolls removed in part 2:", removed)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
