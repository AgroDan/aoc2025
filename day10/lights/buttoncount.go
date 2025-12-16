package lights

import (
	"fmt"
	"utils"
)

type Button struct {
	Indices []int
	Name    string
}

/*
	So here's the plan:
	- Start with target numbers
	- Start with button 0
	- Try pressing it up to as many times as is possible without exceeding any target
	- subtract this effect from the target
		- use DFS this time, aka using a stack instead of a queue
	- move to next button, repeat
	- if at the end all the numbers are zero, we got one!
	- Find all solutions, pick the smallest one.
*/

func (j *JoltMachine) SolveCountsDFS() (int, error) {

	// target will be the "remaining" state that we'll subtract from
	// This will be _how many more we need to achieve the target
	// at index i_
	remaining := make([]int, len(j.Joltage))
	copy(remaining, j.Joltage)

	bestCounts := make([]int, len(j.Buttons))
	currCounts := make([]int, len(j.Buttons))

	// Now do a dumb greedy approximation, instead of choosing infinity
	// or some gigantic number. The absolute max total presses will be
	// the sum of all targets.
	upperBound := utils.SumIntSlice(j.Joltage)
	found := false

	// precompute max "fanout" of button
	// ie how many indices it can increment in one press
	maxFanout := 0
	for _, btn := range j.Buttons {
		if len(btn) > maxFanout {
			maxFanout = len(btn)
		}
	}
	if maxFanout == 0 {
		return -1, fmt.Errorf("no buttons defined")
	}
	// get lowest target to help prune search
	// absolute least amount of presses is the highest number in the
	// list of targets
	// lowerBound := utils.MaxIntSlice(j.Joltage)

	// a little lambda func my friends?
	var dfs func(i, total int)
	dfs = func(i, total int) {
		// Global lower-bound pruning, because all the cool kids are using Z3
		// compute a better lower-bound on how many presses are still needed:
		maxRem := utils.MaxIntSlice(remaining)
		if maxRem < 0 {
			// just in case
			return
		}

		sumRem := utils.SumIntSlice(remaining)

		// basic lower bound, at least maxRem presses are needed
		lowerBound := maxRem

		// cuz I'm extra
		// still need ceiling(sumrem/maxfanout)
		if sumRem > 0 {
			lb := (sumRem + maxFanout - 1) / maxFanout
			if lb > lowerBound {
				lowerBound = lb
			}
		}

		// even if BEST-case future presses can't beat our current best,
		// shoot it in the face
		if total+lowerBound >= upperBound {
			return
		}

		// Capacity pruning: for each index, check if remaining buttons can possibly
		// supply enough increments to satify remaining
		for idx, need := range remaining {
			if need <= 0 {
				continue
			}

			capacity := 0
			// look only at buttons i -> end (still need to be assigned)
			for b := i; b < len(j.Buttons); b++ {
				btnB := j.Buttons[b]

				// compute maxpress for this button under current remaining
				maxPressB := 9999999
				for _, k := range btnB {
					if remaining[k] < maxPressB {
						maxPressB = remaining[k]
					}
				}
				if maxPressB < 0 {
					maxPressB = 0
				}

				// if this button affects index 'idx' it can contribute
				// at most maxPressB to that index
				for _, k := range btnB {
					if k == idx {
						capacity += maxPressB
						break
					}
				}
			}

			// if even the max possible contributions can't reach need,
			// then shoot it in the face with great prejudice
			if capacity < need {
				return
			}
		}

		// maybe put this back?
		// if total+utils.MaxIntSlice(remaining) > upperBound {
		// 	// toss it
		// 	return
		// }

		// if we've gone through all the buttons,
		// check if we have a valid solution.
		if i == len(j.Buttons) {
			for _, v := range remaining {
				if v != 0 {
					return // not a valid solution
				}
			}

			// otherwise, valid solution!
			found = true
			fmt.Printf("Found valid solution with total presses %d\n", total)
			upperBound = total
			copy(bestCounts, currCounts)
			return
		}

		btn := j.Buttons[i]
		// This probably won't happen but just in case
		if len(btn) == 0 {
			currCounts[i] = 0
			dfs(i+1, total)
			return
		}

		// compute max we can press this button
		// without making remaining[idx] negative
		maxPresses := 9999999
		for _, idx := range btn {
			if idx < 0 || idx >= len(remaining) {
				panic(fmt.Sprintf("button index %d out of range", idx))
			}
			if remaining[idx] < maxPresses {
				maxPresses = remaining[idx]
			}
		}
		if maxPresses < 0 {
			maxPresses = 0
		}

		// Now try all possibilities from 0..maxPresses
		for c := 0; c <= maxPresses; c++ {
			currCounts[i] = c

			if c > 0 {
				for _, idx := range btn {
					remaining[idx] -= c
				}
			}

			dfs(i+1, total+c)

			// backtrack
			if c > 0 {
				for _, idx := range btn {
					remaining[idx] += c
				}
			}
		}
	}

	dfs(0, 0)

	if !found {
		return -1, fmt.Errorf("no solution found")
	}

	// print best counts for debugging
	/*
		fmt.Printf("Best button counts:\n")
		for i, c := range bestCounts {
			fmt.Printf("  Button %d: %d presses\n", i, c)
		}
	*/

	return upperBound, nil
}
