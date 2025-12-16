package presents

import (
	"fmt"
	"strconv"
	"strings"
)

// This will define the presents objects

type Present struct {
	index int
	shape [3][3]rune
}

func (p Present) Print() {
	fmt.Printf("Present index: %d\n", p.index)
	fmt.Printf("Shape:\n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%c", p.shape[i][j])
		}
		fmt.Printf("\n")
	}
}

func ParsePresent(presentBlob []string) Present {
	var p Present
	idx := strings.TrimRight(presentBlob[0], ":")
	index, err := strconv.Atoi(idx)
	if err != nil {
		panic(err)
	}
	p.index = index

	var shape [3][3]rune
	var i int = 0
	for _, line := range presentBlob[1:] {
		for j, char := range line {
			shape[i][j] = char
		}
		i++
	}
	p.shape = shape
	return p
}

// Define the challenge region
type Region struct {
	height, width int
	PresentIdx    []int
}

func ParseRegion(line string) Region {
	parts := strings.Split(line, ":")
	dimensions := strings.Split(parts[0], "x")

	// First, height/width
	width, err := strconv.Atoi(dimensions[0])
	if err != nil {
		panic(err)
	}
	height, err := strconv.Atoi(dimensions[1])
	if err != nil {
		panic(err)
	}

	var presentIdx []int
	if len(parts) > 1 {
		// First, clean up any additional spaces
		parts[1] = strings.TrimSpace(parts[1])

		// Now split on spaces
		indices := strings.Split(parts[1], " ")

		// now parse each index
		for _, idxStr := range indices {
			idx, err := strconv.Atoi(strings.TrimSpace(idxStr))
			if err != nil {
				panic(err)
			}
			presentIdx = append(presentIdx, idx)
		}
	}

	return Region{
		height:     height,
		width:      width,
		PresentIdx: presentIdx,
	}
}

func (r Region) Print() {
	fmt.Printf("Region width: %d height: %d\n", r.width, r.height)
	fmt.Printf("Present indices: %v\n", r.PresentIdx)
}

func (p Present) Volume() int {
	// Just returns the volume based on the shape
	volume := 0
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			if p.shape[y][x] == '#' {
				volume++
			}
		}
	}
	return volume
}

func (r Region) Area() int {
	return r.width * r.height
}
