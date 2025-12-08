package main

import (
	"day8/junctions"
	"flag"
	"fmt"
	"time"
	"utils"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")
	wireAmt := flag.Int("w", 1000, "Amount of wire to use for circuits")
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

	j, err := junctions.JunctionList(lines)
	if err != nil {
		panic(err)
	}

	// // Link them first...
	// junctions.LinkJunctions(j)

	first, second, third := junctions.Top3LongestCircuitsWithWire(j, *wireAmt)

	// // Now get the top 3 longest links
	// first, second, third := junctions.Top3LongestLinks(j)

	fmt.Printf("Answer to part 1: %d\n", first*second*third)

	juncA, juncB, err := junctions.ConnectAllJunctions(j)
	if err != nil {
		fmt.Println("Error connecting all junctions:", err)
	} else {
		fmt.Printf("Last junctions connected: %v and %v\n", juncA, juncB)
		fmt.Printf("Answer to question 2: %d\n", juncA.X*juncB.X)
	}

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
