package rectangles

import (
	"slices"
	"utils"
)

// This, knowing now that the points are ordered in such a way that each
// successive point is either directly horizontal or vertical from the
// last, will allow me to draw lines between them to create one giant
// amorphous blob. I can then use a flood fill algorithm to find every
// single point inside that blob, and use it as a map so I can determine
// every point that is considered inside the rectangle. I can then use it
// to eliminate any rectangles that contain points that _aren't_ inside
// the blob.

func CreateBorder(pts []utils.Coord) *utils.GSet[utils.Coord] {
	borderSet := utils.NewGSet[utils.Coord]()

	for i := 0; i < len(pts); i++ {
		from, to := pts[i], pts[(i+1)%len(pts)]

		// Determine if horizontal or vertical line
		if from.X == to.X {
			// Vertical line
			if from.Y < to.Y {
				for y := from.Y; y <= to.Y; y++ {
					borderSet.Add(utils.Coord{X: from.X, Y: y})
				}
			} else {
				for y := to.Y; y <= from.Y; y++ {
					borderSet.Add(utils.Coord{X: from.X, Y: y})
				}
			}
		} else if from.Y == to.Y {
			// Horizontal line
			if from.X < to.X {
				for x := from.X; x <= to.X; x++ {
					borderSet.Add(utils.Coord{X: x, Y: from.Y})
				}
			} else {
				for x := to.X; x <= from.X; x++ {
					borderSet.Add(utils.Coord{X: x, Y: from.Y})
				}
			}
		} else {
			panic("Non-straight line detected between points!")
		}
	}
	return borderSet
}

// Now to determine if a point is inside the border using a flood fill
// algorithm. I'll start at a point I know should be inside the border.
// In this case, I'll just use the average of all the points given,
// then flood fill outwards. If I hit the map edge, then it's not enclosed
// so I should try another point that isn't included in the flood that
// I created (I'll be using a finite set).

func FloodFill(border *utils.GSet[utils.Coord], start utils.Coord, maxX, maxY int) (*utils.GSet[utils.Coord], bool) {
	filled := utils.NewGSet[utils.Coord]()
	toFill := []utils.Coord{start}
	enclosed := true

	neighbors := start.AllAvailable()
	toFill = append(toFill, neighbors...)
	for len(toFill) > 0 {
		current := toFill[0]
		toFill = toFill[1:]

		// If already filled, skip
		if filled.Contains(current) {
			continue
		}

		// If on border, skip
		if border.Contains(current) {
			continue
		}

		// If out of bounds, then not enclosed
		if current.X < 0 || current.X > maxX || current.Y < 0 || current.Y > maxY {
			enclosed = false
			continue
		}

		// Mark as filled
		filled.Add(current)

		// Add neighbors to fill list
		for _, n := range current.AllAvailable() {
			if !filled.Contains(n) && !border.Contains(n) {
				toFill = append(toFill, n)
			}
		}
	}

	return filled, enclosed
}

func SetInside(border *utils.GSet[utils.Coord]) *utils.GSet[utils.Coord] {
	// We'll determine the external border, and we'll start with the first point
	// and check all neighbors to find where is the inside.

	borderCoords := border.ToSlice()

	// First, get the max X and Y values. Handy slices.MaxFunc function!
	maxX := slices.MaxFunc(borderCoords, func(a, b utils.Coord) int {
		return a.X - b.X
	}).X
	maxY := slices.MaxFunc(borderCoords, func(a, b utils.Coord) int {
		return a.Y - b.Y
	}).Y

	var insideSet *utils.GSet[utils.Coord]

	neighbors := borderCoords[0].Neighbors()
	for _, n := range neighbors {
		if border.Contains(n) {
			continue
		}
		checkSet, enclosed := FloodFill(border, n, maxX, maxY)
		if enclosed {
			insideSet = checkSet
			break
		}
	}

	// Now we have the inside data and the border, so return a combined set
	return border.Union(insideSet)
}

// and now, finally, given two coordinates, this will check every point in that
// coordinate rectangle to see if enclosed.

func IsRectangleEnclosed(from, to utils.Coord, insideSet *utils.GSet[utils.Coord]) bool {
	// this should check for a full rectangle defined by a and b
	// so get the min and max X and Y values will tell me where
	// to start drawing the rectangle. I'll start from the smallest
	// X and Y values to the largest, so that should be from top left
	// to bottom right
	var a utils.Coord
	var b utils.Coord
	// fmt.Printf("Debug: Checking rectangle from %+v to %+v\n", from, to)

	if from.X < to.X {
		a.X = from.X
		b.X = to.X
	} else {
		a.X = to.X
		b.X = from.X
	}

	if from.Y < to.Y {
		a.Y = from.Y
		b.Y = to.Y
	} else {
		a.Y = to.Y
		b.Y = from.Y
	}

	for x := a.X; x <= b.X; x++ {
		for y := a.Y; y <= b.Y; y++ {
			checkPoint := utils.Coord{X: x, Y: y}
			if !insideSet.Contains(checkPoint) {
				return false
			}
		}
	}
	return true
}

// GetPerimeter will return the perimeter of a rectangle given two coords
// the perimeter will be a slice of coordinates representing the border
func GetPerimeter(a, b utils.Coord) []utils.Coord {
	var topLeft, bottomRight utils.Coord

	if a.X < b.X {
		topLeft.X = a.X
		bottomRight.X = b.X
	} else {
		topLeft.X = b.X
		bottomRight.X = a.X
	}

	if a.Y < b.Y {
		topLeft.Y = a.Y
		bottomRight.Y = b.Y
	} else {
		topLeft.Y = b.Y
		bottomRight.Y = a.Y
	}

	perimeter := make([]utils.Coord, 0)

	// Top edge
	for x := topLeft.X; x <= bottomRight.X; x++ {
		perimeter = append(perimeter, utils.Coord{X: x, Y: topLeft.Y})
	}
	// Bottom edge
	for x := topLeft.X; x <= bottomRight.X; x++ {
		perimeter = append(perimeter, utils.Coord{X: x, Y: bottomRight.Y})
	}
	// Left edge
	for y := topLeft.Y; y <= bottomRight.Y; y++ {
		perimeter = append(perimeter, utils.Coord{X: topLeft.X, Y: y})
	}
	// Right edge
	for y := topLeft.Y; y <= bottomRight.Y; y++ {
		perimeter = append(perimeter, utils.Coord{X: bottomRight.X, Y: y})
	}

	return perimeter
}
