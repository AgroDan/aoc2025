package main

import (
	"day6/cmath"
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
	homework := cmath.ReadHomework(lines)
	result := cmath.DoHomework(homework)
	fmt.Printf("The result of the homework is: %d\n", result)

	// Now part 2...
	homeworkPartTwo := cmath.PlotHomework(lines)
	problems := cmath.GetProblems(homeworkPartTwo)
	resultPartTwo := 0
	for _, p := range problems {
		p.Print()
		solution := p.Solve()
		fmt.Printf("Solves to: %d\n\n", solution)
		resultPartTwo += solution
	}
	fmt.Printf("Sum of all results for part two: %d\n", resultPartTwo)
	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
