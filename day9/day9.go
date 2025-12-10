package main

import (
	"day9/rectangles"
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

	pts := rectangles.ParseInput(lines)
	// runemap := rectangles.CreateRuneMap(pts)
	// runemap.Print()

	// Get the pairs
	pairs := rectangles.GeneratePairs(pts)

	// largest rectangle should be the last item in the list
	// because GeneratePairs also sorts by manhattan distance
	partOne := rectangles.RectangleArea(pairs[len(pairs)-1].A, pairs[len(pairs)-1].B)
	// fmt.Printf("Largest pair: %+v and %+v with distance %d\n", pairs[len(pairs)-1].A, pairs[len(pairs)-1].B, pairs[len(pairs)-1].Dist)
	fmt.Printf("Part One: Largest rectangle area is %d\n", partOne)

	// Now to do part two, so find the border first...
	fmt.Printf("Creating border...\n")
	borderSet := rectangles.CreateBorder(pts)

	// create a set of all possible points...
	// fmt.Printf("Creating flood fill area...\n")
	// area := rectangles.SetInside(borderSet)

	// Now loop through the pairs backwards until I find the biggest rectangle
	// that is considered all inside the border
	var partTwo int = 0
	fmt.Printf("Total pairs: %d\n", len(pairs))
	provenIn := utils.NewGSet[utils.Coord]()
	for i := len(pairs) - 1; i >= 0; i-- {
		a, b := pairs[i].A, pairs[i].B
		// fmt.Printf("Checking pair: %+v and %+v with area %d\n", a, b, pairs[i].Dist)
		if rectangles.CheckRectangle(a, b, borderSet, provenIn) {
			partTwo = rectangles.RectangleArea(a, b)
			break
		}
	}

	// rectangles.IncludeFloodFill(&runemap, area, 'O')
	// fmt.Println("Runemap with enclosed area filled:")
	// runemap.Print()
	fmt.Printf("Part Two: Largest enclosed rectangle area is %d\n", partTwo)
	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
