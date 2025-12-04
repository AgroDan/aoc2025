package rolls

import "utils"

// This package will work on the paper roll factory floor and do all
// the heavy lifting of finding everything needed.

func FindAccessibleRolls(r utils.Runemap) []utils.Coord {
	retval := []utils.Coord{}

	rolls := r.FindAll('@')
	for _, roll := range rolls {
		// check all eight directions and see how many have rolls already
		neighbors := roll.Neighbors()
		totalRolls := 0
		for _, n := range neighbors {
			if r.IsInBounds(n) {
				char, _ := r.Get(n)
				if char == '@' {
					totalRolls++
				}
			}
		}
		if totalRolls <= 3 {
			retval = append(retval, roll)
		}
	}

	return retval
}

// Now for part two. Let's start removing rolls. Luckily I have functions
// that can help me out with this.

func RemoveRolls(r *utils.Runemap) int {
	removedCount := 0
	for {
		accessibleRolls := FindAccessibleRolls(*r)
		if len(accessibleRolls) == 0 {
			break
		}
		for _, roll := range accessibleRolls {
			r.Set(roll, '.')
			removedCount++
		}
	}
	return removedCount
}
