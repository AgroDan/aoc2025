package tachyon

import (
	"fmt"
	"utils"
)

// So now the puzzle asks for how many possible routes a single tachyon
// beam can possibly take throughout the manifold. Whenever a beam hits
// a splitter, it will create two possible directions _from that point_.
// However, if it hits another splitter, it could branch off into two
// different directions from there. This could branch off into a
// huge amount so I think the answer here is memoization. This basic
// idea was so similar to last year's "robot pushing buttons to control
// other robots down the line" challenge that I feel it's probably the
// best and only way to figure this out. So let's try it.

// I'll create a function that will compute how many paths are possible
// from a given coordinate, assuming that coordinate is a splitter.

func CountPathsFromSplitter(m utils.Runemap, coord utils.Coord, cache *utils.Cache) int {
	// This function is meant to run from the first splitter found.
	paths := 0

	// First, check SW
	swDir := coord.Peek(utils.SW)

	// Keep going down until we hit the bottom
	for {
		downDir := swDir.Peek(utils.S)
		downRune, err := m.Get(downDir)
		if err != nil {
			// Out of bounds, this is one single path
			paths++
			break
		}

		if downRune == '^' {
			cacheKey := fmt.Sprintf("paths_%d_%d", downDir.X, downDir.Y)
			subPaths := cache.Get(cacheKey, func() interface{} {
				return CountPathsFromSplitter(m, downDir, cache)
			}).(int)

			paths += subPaths
			break
		}

		// Move down
		swDir = downDir
	}

	// Now check SE
	seDir := coord.Peek(utils.SE)

	// Keep going down until we hit the bottom
	for {
		downDir := seDir.Peek(utils.S)
		downRune, err := m.Get(downDir)
		if err != nil {
			// Out of bounds, this is one single path
			paths++
			break
		}

		if downRune == '^' {
			cacheKey := fmt.Sprintf("paths_%d_%d", downDir.X, downDir.Y)
			subPaths := cache.Get(cacheKey, func() interface{} {
				return CountPathsFromSplitter(m, downDir, cache)
			}).(int)

			paths += subPaths
			break
		}

		// Move down
		seDir = downDir
	}

	return paths
}

// Now, given the map and the start coord, we'll find the first splitter and go from there.
func TotalTachyonPaths(m utils.Runemap, start utils.Coord) int {
	// keep going until we find the first splitter
	for {
		downDir := start.Peek(utils.S)
		downRune, err := m.Get(downDir)
		if err != nil {
			// Out of bounds, no splitters found
			return 0
		}

		if downRune == '^' {
			// Found the first splitter, now to count paths from here
			cache := utils.NewCache()
			return CountPathsFromSplitter(m, downDir, cache)
		}
		// Move down
		start = downDir
	}
}
