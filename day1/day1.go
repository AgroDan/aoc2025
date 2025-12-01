package main

import (
	"day1/dial"
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

	d := dial.NewDial(50)
	totalTimesLandedAtZero := 0
	for _, line := range lines {
		d.ParseInstruction(line)
		if d.GetPosition() == 0 {
			totalTimesLandedAtZero++
		}
	}
	fmt.Printf("Total rotations that landed at 0 for part 1: %d\n", totalTimesLandedAtZero)

	d2 := dial.NewDial(50)
	totalRotations := 0
	for _, line := range lines {
		totalRotations += d2.ParseInstructionCountRotations(line)
	}
	fmt.Printf("Total rotations that passed 0 for part 2: %d\n", totalRotations)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
