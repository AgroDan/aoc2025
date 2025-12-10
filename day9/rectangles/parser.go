package rectangles

import (
	"fmt"
	"utils"
)

// This will ingest the data line by line as coordinates and return
// a slice of coordinates.
func ParseInput(input []string) []utils.Coord {
	var coords []utils.Coord
	for _, line := range input {
		var x, y int
		_, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if err != nil {
			panic("failed to parse input line: " + line)
		}
		coords = append(coords, utils.Coord{X: x, Y: y})
	}
	return coords
}

// Now a helper to find the max X and Y values from the list of coords
func FindMaxCoords(coords []utils.Coord) (int, int) {
	maxX, maxY := 0, 0
	for _, coord := range coords {
		if coord.X > maxX {
			maxX = coord.X
		}
		if coord.Y > maxY {
			maxY = coord.Y
		}
	}
	return maxX, maxY
}

// Now let's create a runemap with the given dimensions
func CreateRuneMap(coords []utils.Coord) utils.Runemap {
	mapSeed := make([]string, 0)
	maxX, maxY := FindMaxCoords(coords)
	maxX += 1 // just to pad it
	maxY += 1
	for y := 0; y <= maxY; y++ {
		row := ""
		for x := 0; x <= maxX; x++ {
			row += "."
		}
		mapSeed = append(mapSeed, row)
	}

	retMap := utils.NewRunemap(mapSeed)

	// Now set the points
	for _, coord := range coords {
		retMap.Set(coord, '#')
	}
	return retMap
}

func IncludeFloodFill(rm *utils.Runemap, fillSet *utils.GSet[utils.Coord], fillChar rune) {
	fillSlice := fillSet.ToSlice()
	for _, coord := range fillSlice {
		rm.Set(coord, fillChar)
	}
}
