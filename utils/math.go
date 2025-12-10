package utils

func EuclideanModulus(x, y int) int {
	// This is the euclidean modulus, which is actually different
	// from Go's default modulus operation, which is actually
	// "technically correct." In this case, -1 % 10 will equal
	// 9, where in Go -1 % 10 will equal -1. This is useful for
	// things like rotating around a map where if you go over the
	// map border, it will teleport you to the opposite side.
	rem := x % y
	if rem < 0 {
		rem += y
	}

	return rem
}

// helper for getting the absolute value so I don't have to do dumbass
// int(float64()-float64()) conversions
//
// this used to be in coords.go but moved here for refactoring, makes
// more sense to be in math.go I think
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(a, b int) int {
	// already done but this is just easier since it's ints and not floats
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
