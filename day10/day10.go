package main

import (
	"day10/lights"
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

	lightSet := make([]*lights.Machine, 0)

	partonetotal := 0
	for i, line := range lines {
		lightSet = append(lightSet, lights.NewMachine(line))
		// lightSet[i].Print()
		solved, err := lightSet[i].Solve()
		if err != nil {
			fmt.Println("Error solving machine:", err)
		} else {
			// fmt.Printf("Solved with %d steps\n", solved)
			partonetotal += solved
		}
		// fmt.Printf("------\n")
	}
	fmt.Printf("Part One: Total steps to solve all machines: %d\n", partonetotal)

	// Have to re-parse for part 2 since we're doing it a little bit differently
	parttwototal := 0
	joltSet := make([]*lights.JoltMachine, 0)
	for i, line := range lines {
		joltSet = append(joltSet, lights.NewJoltMachine(line))
		// joltSet[i].Print()
		presses := lights.PartTwoSolver(joltSet[i].Buttons, joltSet[i].Joltage)
		// if err != nil {
		// 	fmt.Println("Jolt machine unsolveable:", err)
		// } else {
		// 	// fmt.Printf("Solved jolt machine with %d steps\n", solved)
		// 	parttwototal += thisLightCount
		// }
		// fmt.Printf("Best Presses that work: %d\n", presses)
		parttwototal += presses
	}

	fmt.Printf("Part Two: Total steps to solve all jolt machines: %d\n", parttwototal)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
