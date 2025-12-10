package rectangles

import (
	"fmt"
	"slices"
	"utils"
)

// For this one, just like in Day 8, I'll generate a set of pairs
// and calculate the Manhattan distance between them. This should
// give me the largest rectangle.

type Pair struct {
	A, B utils.Coord
	Dist int
}

func (p Pair) String() string {
	return fmt.Sprintf("Pair(%+v, %+v) Dist: %d", p.A, p.B, p.Dist)
}

func GeneratePairs(coords []utils.Coord) []Pair {
	pairs := make([]Pair, 0)
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			// Instead of getting the manhattan distance, I'll just get the area
			// dist := utils.ManhattanDistance(coords[i], coords[j])
			dist := RectangleArea(coords[i], coords[j])
			pairs = append(pairs, Pair{A: coords[i], B: coords[j], Dist: dist})
		}
	}

	// Now sort them by distance
	slices.SortFunc(pairs, func(a, b Pair) int {
		return a.Dist - b.Dist
	})
	return pairs
}

// Also will need to find the Area of a rectangle given two coords,
// as per the puzzle request
func RectangleArea(a, b utils.Coord) int {
	width := utils.Abs(a.X-b.X) + 1 // have to account for the full border
	height := utils.Abs(a.Y-b.Y) + 1
	// fmt.Printf("Debug: Width %d, Height %d\n", width, height)
	return width * height
}
