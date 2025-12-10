package rectangles

import (
	"fmt"
	"utils"
)

// OK! This is a MUCH more efficient method of determining if a point is inside
// the giant amorphous blob of red and green than doing a flood fill for the entire
// surface area of the rectangle. I know I did this in a previous AOC but I totally
// forgot about it until I looked for a better way of determining if a point is inside
// of a polygon or not. Enter the Raycasting algorithm!

// basically, pick any point outside the polygon and draw a line from it to the border.
// However many times it crosses the border with ALTERNATING border crossings...if it's
// odd, then we are inside the polygon. If even, we are outside.

// The "alternating border crossings" part is tricky, but it helps for a grid like this:

// ....#.......
// ...#.#......
// ..#...#.....
// .#.###..#...
// ####.#.X.#..
// .....#...#..
// .#####...#..
// .#......##..
// ..######....

// If we choose to draw a line from X to the left, it crosses 5 border points, but
// in reality, we only cross two. Let's count it until we hit a non-border point as one.

func IsPointInsidePolygonWest(border *utils.GSet[utils.Coord], point utils.Coord) bool {
	// first of all, am I on the border?
	if border.Contains(point) {
		return true
	}

	// Start from a point outside the polygon, say (-1, point.Y)
	outsidePoint := utils.Coord{X: -1, Y: point.Y}

	// Draw a line from outsidePoint to point
	// Since we are only moving horizontally, we can just iterate over X values
	crossings := 0

	// if we're on the border, we're technically inside
	if border.Contains(point) {
		return true
	}

	for x := outsidePoint.X; x <= point.X; x++ {
		currentCoord := utils.Coord{X: x, Y: point.Y}
		if border.Contains(currentCoord) {
			// We hit a border point, now count consecutive border points as one crossing
			crossings++
			// Move x forward until we hit a non-border point
			innerX := x + 1
			for innerX <= point.X && border.Contains(utils.Coord{X: innerX, Y: point.Y}) {
				innerX++
			}
			x = innerX - 1 // adjust x to the last border point
		}
	}

	// If crossings is odd, point is inside; if even, point is outside
	return crossings%2 == 1
}

func IsPointInsidePolygonNorth(border *utils.GSet[utils.Coord], point utils.Coord) bool {
	// am i on the border?
	if border.Contains(point) {
		return true
	}

	// Start from a point outside the polygon, say (point.X, -1)
	outsidePoint := utils.Coord{X: point.X, Y: -1}

	// Draw a line from outsidePoint to point
	// Since we are only moving vertically, we can just iterate over Y values
	crossings := 0

	// if we're on the border, we're technically inside
	if border.Contains(point) {
		return true
	}

	for y := outsidePoint.Y; y <= point.Y; y++ {
		currentCoord := utils.Coord{X: point.X, Y: y}
		if border.Contains(currentCoord) {
			// We hit a border point, now count consecutive border points as one crossing
			crossings++
			// Move y forward until we hit a non-border point
			innerY := y + 1
			for innerY <= point.Y && border.Contains(utils.Coord{X: point.X, Y: innerY}) {
				innerY++
			}
			y = innerY - 1 // adjust y to the last border point
		}
	}

	// If crossings is odd, point is inside; if even, point is outside
	return crossings%2 == 1
}

// Helper for raytracing
func IsRectangleEnclosedRT(from, to utils.Coord, borderSet *utils.GSet[utils.Coord]) bool {
	// this should check for a full rectangle defined by a and b
	// so get the min and max X and Y values will tell me where
	// to start drawing the rectangle. I'll start from the smallest
	// X and Y values to the largest, so that should be from top left
	// to bottom right

	perim := GetPerimeter(from, to)
	fmt.Printf("Debug: Rectangle perimeter has %d points\n", len(perim))

	for _, pt := range perim {
		if !IsPointInsidePolygonWest(borderSet, pt) {
			return false
		}
	}
	return true
}

// Now what I'm going to do is traverse the border coordinates. If a coordinate is NOT
// a border coordinate that I'll pass to it as a parameter of the larger polygon, then
// I'll perform a raycast from that point. If it's fully enclosed then we're good to go.

func IsRectangleEnclosedRT2(from, to utils.Coord, borderSet *utils.GSet[utils.Coord], ptCache *utils.Cache) bool {
	maxX := utils.Max(from.X, to.X)
	minX := utils.Min(from.X, to.X)
	maxY := utils.Max(from.Y, to.Y)
	minY := utils.Min(from.Y, to.Y)

	// First, loop through the horizontal lines
	for x := minX; x <= maxX; x++ {
		topPoint := utils.Coord{X: x, Y: minY}
		bottomPoint := utils.Coord{X: x, Y: maxY}
		if !borderSet.Contains(topPoint) && !IsPointInsidePolygonNorth(borderSet, topPoint) {
			return false
		}
		if !IsPointInsidePolygonNorth(borderSet, bottomPoint) {
			return false
		}
	}

	// Now the vertical lines
	for y := minY; y <= maxY; y++ {
		leftPoint := utils.Coord{X: minX, Y: y}
		rightPoint := utils.Coord{X: maxX, Y: y}

		if !borderSet.Contains(leftPoint) && !IsPointInsidePolygonWest(borderSet, leftPoint) {
			return false
		}
		if !IsPointInsidePolygonWest(borderSet, rightPoint) {
			return false
		}
	}

	return true
}
