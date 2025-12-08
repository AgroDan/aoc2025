package junctions

import (
	"fmt"
	"math"
)

// This will define the object and get the Euclidian distance between it and
// another junction. It will also have the ability to link to another junction
// as they will be connected via christmas lights

type Junction struct {
	X, Y, Z int
	Link    *Junction
}

func NewJunction(x, y, z int) *Junction {
	return &Junction{
		X:    x,
		Y:    y,
		Z:    z,
		Link: nil,
	}
}

func (j *Junction) String() string {
	return fmt.Sprintf("X: %d, Y: %d, Z: %d, IsLinked: %t", j.X, j.Y, j.Z, j.Link != nil)
}

func (j *Junction) Distance(other *Junction) float64 {
	// This is the distance according to wikipedia...
	// https://en.wikipedia.org/wiki/Euclidean_distance
	deltax := math.Pow(float64(j.X-other.X), 2)
	deltay := math.Pow(float64(j.Y-other.Y), 2)
	deltaz := math.Pow(float64(j.Z-other.Z), 2)
	return math.Sqrt(deltax + deltay + deltaz)
}
