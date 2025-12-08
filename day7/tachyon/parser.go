package tachyon

import "utils"

type Beam struct {
	pos utils.Coord
}

func NewManifold(lines []string) (utils.Runemap, utils.Coord) {
	m := utils.NewRunemap(lines)
	start, _ := m.Find('S')
	return m, start
}

// For this challenge, create queue that allows for more tachyon beams to split.
// But then we also have to make sure we search for when two beams converge into
// one, so maybe get a breadcrumb trail, and when a beam enters a previously
// traversed path, just count it once.

// Maybe use a finite set for that? A map of Coordinates?

// I dunno, i'm exhausted. Going to bed.

// ...and back. So now I'm going to start at the starting point, move down
// and add that movement into a queue. If it hits a splitter, it adds two
// more tachyon beams into the queue if and ONLY if those paths have not
// been traversed yet. If any beam ever splits, add it it to the total
// count. This is the puzzle answer!

func TraverseBeams(m utils.Runemap, start utils.Coord) (int, map[utils.Coord]struct{}) {
	// We're going to use the utils.coord object to determine directions.
	totalSplits := 0

	// Let's set up the queue
	beamQueue := utils.NewGQueue[utils.Coord]()
	beamQueue.Enqueue(start)

	// now to set up a finite set of traversed paths
	var exists = struct{}{}
	traversed := make(map[utils.Coord]struct{})
	traversed[start] = exists

	// Now to process the queue
	for !beamQueue.IsEmpty() {
		current, _ := beamQueue.Dequeue()

		// First, check down.
		downDir := current.Peek(utils.S)
		downRune, err := m.Get(downDir)
		if err != nil {
			// Out of bounds, ignore
			continue
		}

		// Now first, check if we've been there.
		if _, found := traversed[downDir]; found {
			// Already been here, this beam has re-integrated
			// to another beam
			continue
		}

		// Now check to see if there's a splitter
		if downRune == '^' {
			// Splitter found, first increment the splitter count
			totalSplits++

			// Now let's check SE and SW, same rules apply if we've been there
			seDir := current.Peek(utils.SE)
			// seRune, err := m.Get(seDir)
			// if err != nil {
			// 	// out of bounds, probably won't ever happen but just in case
			// 	// and if it does then swRune would error out too right?
			// 	// REMEMBER THIS IF IT ERRORS OUT!
			// 	continue
			// }

			// Now check if we've been there...
			if _, found := traversed[seDir]; !found {
				// haven't been here, this is valid.
				beamQueue.Enqueue(seDir)
				traversed[seDir] = exists
			}

			// Now for swDir
			swDir := current.Peek(utils.SW)
			// swRune, err := m.Get(swDir)
			// if err != nil {
			// 	// out of bounds, skip
			// 	continue
			// }

			// Now check if we've been there...
			if _, found := traversed[swDir]; !found {
				// haven't been here, this is valid.
				beamQueue.Enqueue(swDir)
				traversed[swDir] = exists
			}

			// now continue so we don't blow through this splitter
			continue
		}

		// Otherwise just move down normally.
		beamQueue.Enqueue(downDir)
		traversed[downDir] = exists
	}

	return totalSplits, traversed
}
