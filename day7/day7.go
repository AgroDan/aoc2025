package main

import (
	"day7/tachyon"
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
	tachyonMap, start := tachyon.NewManifold(lines)
	tachyonMap.Print()
	splits, traversed := tachyon.TraverseBeams(tachyonMap, start)
	// splits, _ := tachyon.TraverseBeams(tachyonMap, start)

	for coord, _ := range traversed {
		tachyonMap.Set(coord, '|')
	}

	fmt.Println("Traversed Map:")
	tachyonMap.Print()
	fmt.Printf("Total tachyon splits: %d\n", splits)

	totalPaths := tachyon.TotalTachyonPaths(tachyonMap, start)
	fmt.Printf("Total quantum tachyon paths: %d\n", totalPaths)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
