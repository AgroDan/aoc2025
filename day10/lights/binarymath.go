package lights

import (
	"slices"
)

// THIS is the key here:
// https://www.reddit.com/r/adventofcode/comments/1pk87hl/2025_day_10_part_2_bifurcate_your_way_to_victory/
//
// So...after some reading up on people far smarter than me using tricks far cooler
// than anything I've tried, I read up on the concept of using binary math to use
// _parity_ to determine the possible outcomes of button presses. The idea is that
// if you have a joltage target and a set of buttons that can increment certain
// joltage indices by 1 each time they're pressed, you can potentially determine
// the combination of buttons that can achieve the target by determining which
// buttons would produce the proper outcome based on parity alone! Because this
// is one of those things that I just need to code out the actual logic for by
// writing helper functions, I'm going to just add as much context as I can in bylines

// This is going to be similar to my part 1 deduction, using a mask to get combinations
// and checking them against the outcome

func allZeroes(arr []int) bool {
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}

func allEvenParity(arr []int) bool {
	for _, v := range arr {
		if v%2 != 0 {
			return false
		}
	}
	return true
}

func howManyButtonPresses(mask int) int {
	// basically counts how many bits are set to 1
	// in the mask and returns that count. This is
	// a helper function super-specific to this particular
	// challenge so this doesn't get exported
	count := 0
	for i := 0; i < 64; i++ {
		if (mask & (1 << i)) != 0 {
			count++
		}
	}
	return count
}

func halfJoltageValues(arr []int) []int {
	// you better make sure all values are even!
	halved := make([]int, len(arr))
	for i, v := range arr {
		halved[i] = v / 2
	}
	return halved
}

func getParityValues(arr []int) []int {
	parity := make([]int, len(arr))
	for i, v := range arr {
		if v%2 != 0 {
			parity[i] = 1
		} else {
			parity[i] = 0
		}
	}
	return parity
}

func anyNegative(arr []int) bool {
	// any numbers in this slice negative?
	for _, v := range arr {
		if v < 0 {
			return true
		}
	}
	return false
}

// just better than all 9's i think
const INF = int(^uint64(0) >> 1)

func PartTwoSolver(btnArr [][]int, joltages []int) int {
	// This will be a recursive function defined by the above link
	// also this can probably be cached right?
	// also i got rid of the btnArr []uint64 parameter and changed it
	// to a regular integer because just dealing with endianness
	// was making me question my life choices

	// First of all, if joltage slice is all zeroes, no presses
	// are needed so return what we have already
	if allZeroes(joltages) {
		// fmt.Printf("Found solution with %d presses\n", presses)
		return 0
	}

	// Now get the target joltage affect we want to find, and we
	// use the parity of each button to determine the state.
	targetParity := getParityValues(joltages)

	// Now build the amount of possible combinations. The amount
	// of buttons is len(btnArr), so total combinations is 2^n
	n := len(btnArr)
	totalCombos := 1 << n // also 2^n! bitwise FTW

	// just going to return the lowest number here
	bestPresses := INF

	// Now each number in this range will represent the combination
	// of buttons pressed, represented in binary. So for example,
	// if we have 3 buttons and we want to show the combination
	// of pressing button 0 and 2, it would be 0b101 in binary!
	// But really we're not that granular about it, we're just
	// going to use a number to represent a specific combination
	// of pressing one or more buttons. Or none!
	for mask := 0; mask < totalCombos; mask++ {
		// skipping 0 because it doesn't press anything
		// fmt.Printf("Testing combination mask: %0*b\n", n, mask)
		// For each combination, we need to determine the parity
		// effect of pressing those buttons

		// This is the temp joltage array that we'll modify by
		// pressing a combination of buttons
		testJoltageValues := make([]int, len(joltages))
		copy(testJoltageValues, joltages)

		// deltaParity = delta mode 2 (0/1) to match targetParity

		deltaParity := make([]int, len(joltages))

		// n == number of buttons total
		for i := 0; i < n; i++ {
			// in this combination, press the specific button
			// denoted in the maskof this combination set
			if (mask & (1 << i)) != 0 {
				// This button is pressed, so apply its effect
				// fmt.Printf("  pressing button %d\n", i)
				for _, j := range btnArr[i] {
					testJoltageValues[j]-- // subtract delta (button press)
					deltaParity[j] ^= 1    // toggle parity
				}
			}
		}
		// fmt.Printf("  Resulting joltages: %v\n", testJoltageValues)

		// Any negative values? If so skip this combo
		if anyNegative(testJoltageValues) {
			continue
		}

		// Is this a target parity?
		if !slices.Equal(deltaParity, targetParity) {
			// Not a match, skip
			continue
		}

		if !allEvenParity(testJoltageValues) {
			continue
			// this is a bad combo so skip
		}

		halvedJoltages := halfJoltageValues(testJoltageValues)

		rest := PartTwoSolver(btnArr, halvedJoltages)
		if rest == INF {
			// impossible so keep going
			continue
		}

		// if we got here, this is a working combo so submit this as a candidate
		candidate := howManyButtonPresses(mask) + (2 * rest)

		if candidate < bestPresses {
			bestPresses = candidate
		}
	}
	return bestPresses
}
