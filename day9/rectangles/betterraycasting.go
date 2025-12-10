package rectangles

import "utils"

// I'm going to see if I can simplify this a bit...

func PointInPolygon(point utils.Coord, border *utils.GSet[utils.Coord], provenIn *utils.GSet[utils.Coord]) bool {
	// Passing the point we are checking, the border set we've already computed,
	// and the provenIn set, which contains all points we already know are inside.

	// First, check if the point is already in the provenIn set
	if provenIn.Contains(point) {
		return true
	}

	// Next, check to see if the point is on the border.
	if border.Contains(point) {
		return true
	}

	// If we got here, ray-cast to the west
	limitCoord := utils.Coord{X: -1, Y: point.Y}
	inside := false
	for x := point.X - 1; x >= limitCoord.X; x-- {
		currentCoord := utils.Coord{X: x, Y: point.Y}
		if border.Contains(currentCoord) {
			inside = !inside
			// Skip consecutive border points
			innerX := x - 1
			for innerX >= limitCoord.X && border.Contains(utils.Coord{X: innerX, Y: point.Y}) {
				innerX--
			}
			x = innerX + 1 // adjust x to the last border point
		}
	}

	// If inside is true, add to provenIn set
	if inside {
		provenIn.Add(point)
	}

	return inside
}

// Now, given two sets, draw a perimeter and check the border for each point
func CheckRectangle(from, to utils.Coord, border *utils.GSet[utils.Coord], provenIn *utils.GSet[utils.Coord]) bool {
	// Draw the rectangle border
	minX := utils.Min(from.X, to.X)
	maxX := utils.Max(from.X, to.X)
	minY := utils.Min(from.Y, to.Y)
	maxY := utils.Max(from.Y, to.Y)
	borderSet := utils.NewGSet[utils.Coord]()
	// Top edge
	for x := minX; x <= maxX; x++ {
		borderSet.Add(utils.Coord{X: x, Y: minY})
	}
	// Bottom edge
	for x := minX; x <= maxX; x++ {
		borderSet.Add(utils.Coord{X: x, Y: maxY})
	}
	// Left edge
	for y := minY; y <= maxY; y++ {
		borderSet.Add(utils.Coord{X: minX, Y: y})
	}
	// Right edge
	for y := minY; y <= maxY; y++ {
		borderSet.Add(utils.Coord{X: maxX, Y: y})
	}

	// Now check each point in borderSet against the main border
	for _, coord := range borderSet.ToSlice() {
		if !PointInPolygon(coord, border, provenIn) {
			return false
		}
	}
	return true
}
