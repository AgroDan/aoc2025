package ranges

import (
	"fmt"
	"strconv"
)

// This package will just be utilities needed to work the challenge

func RepeatedNumbers(num int) []int {
	// This will convert the entire digit into a string
	// and check for repeated patterns within the number.
	// Note that no numbers have trailing zeroes, so since
	// we're getting an integer input we don't have to worry

	invalidIds := []int{}

	numStr := strconv.Itoa(num)
	maxSizeOfPattern := len(numStr) / 2
	// We only care about 2 or more chars, easy win
	if maxSizeOfPattern < 1 {
		return []int{}
	}

	// Now I'm going to create a slice of strings of every possible
	// pattern detected in the numberstring, from 2 to maxSizeOfPattern
	patterns := []string{}
	for size := 2; size <= maxSizeOfPattern; size++ {
		for start := 0; start <= len(numStr)-size; start++ {
			pattern := numStr[start : start+size]
			patterns = append(patterns, pattern)
		}
	}

	// For debugging purposes, let's see how much of a search window we have...
	fmt.Printf("Amount of patterns in %s: %d\n", numStr, len(patterns))

	// FOR FUTURE REFERENCE -- I may need to consider leading zeroes...dunno yet.

	// Now we're going to loop through every single pattern (luckily just incrementally)
	// and check to see if the pattern exists more than once
	for i := 0; i < len(patterns); i++ {
		checkPattern := patterns[i]
		if i == len(patterns)-1 {
			break
		}
		for j := i + 1; j < len(patterns); j++ {
			if checkPattern == patterns[j] {
				numConv, _ := strconv.Atoi(checkPattern)
				invalidIds = append(invalidIds, numConv)
			}
		}
	}
	return invalidIds
}

func RepeatedConsecutiveNumbers(num int) []int {
	// This will convert the entire number into a string, then
	// check for repeated numbers IMMEDIATELY after the first
	// instance only. So this would catch 11, 1122, 123123, etc
	invalidIds := []int{}

	numStr := strconv.Itoa(num)
	maxSizeOfPattern := len(numStr) / 2
	// We only care about 2 or more chars, easy win
	if maxSizeOfPattern < 1 {
		return []int{}
	}

	// Now similar to above, I'm going to create a bunch of patterns
	// starting from single digits and then JUST checking the NEXT
	// pattern to see if it matches. If it does, just toss it in the
	// retval

	for i := 0; i < maxSizeOfPattern; i++ {
		patternSize := i + 1
		for start := 0; start <= len(numStr)-patternSize*2; start++ {
			pattern := numStr[start : start+patternSize]
			nextPattern := numStr[start+patternSize : start+patternSize*2]
			if pattern == nextPattern {
				numConv, _ := strconv.Atoi(pattern + nextPattern)
				invalidIds = append(invalidIds, numConv)
			}
		}
	}

	return invalidIds
}

func FindMirroredNumbers(num int) int {
	// This is embarassing. It's only if the number, split in half, is a
	// mirror image of itself. So I'll just do that. OMG.
	numStr := strconv.Itoa(num)
	if len(numStr)%2 != 0 {
		// only even-sized numbers work here
		return 0
	}

	half := len(numStr) / 2

	if numStr[:half] == numStr[half:] {
		numConv, _ := strconv.Atoi(numStr)
		return numConv
	}
	return 0
}

func FindRepeatedPatterns(num int) int {
	// This will check to see if there is a repeated pattern throughout the _entire_ number.
	// So 1212, 123123, 999999, etc. This will NOT catch 12121, 1231231, etc.
	// returns 0 if the numer is not a repeated pattern.
	numStr := strconv.Itoa(num)
	maxSizeOfPattern := len(numStr) / 2
	if maxSizeOfPattern < 1 {
		return 0
	}

	for size := 1; size <= maxSizeOfPattern; size++ {
		if len(numStr)%size != 0 {
			// This is not a valid pattern size because
			// it wouldn't be repeated completely
			continue
		}
		pattern := numStr[:size]
		repeats := len(numStr) / size
		constructed := ""
		for i := 0; i < repeats; i++ {
			constructed += pattern
		}

		// Basically build a number based on this repeated
		// pattern, and if it matches the original number
		// then the pattern is valid!
		if constructed == numStr {
			numConv, _ := strconv.Atoi(numStr)
			return numConv
		}
	}
	return 0
}
