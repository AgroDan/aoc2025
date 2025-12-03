package batterybank

import (
	"fmt"
	"slices"
	"strconv"
	"utils"
)

type BatteryBank struct {
	Cells []int
}

func NewBatteryBank(cells string) *BatteryBank {
	// This function creates a new battery bank. Reads the
	// line as a string, then converts each single character
	// to the integer value and stores it in the Cells slice.
	// Part two may involve more than one number, so...keep
	// that in mind I guess.
	bb := &BatteryBank{
		Cells: make([]int, len(cells)),
	}
	for i, ch := range cells {
		// I discovered this strange rune to int conversion, by
		// subtracting the rune provided with '0' character, you
		// get the integer value. The reason is because in the ASCII
		// table, the character '0' equates to integer 30, '1' to 31, etc.
		// So if you subtract the rune of an integer, like say 4 (which is 34
		// in ASCII) to '0' (which is 30 in ASCII), you get literal integer of 4!
		// I love this so much
		bb.Cells[i] = int(ch - '0')
	}
	return bb
}

func (bb *BatteryBank) LargestCharge(start, end int) (int, int) {
	// This will search sequentially from the start index and
	// find the largest charge in the battery bank. Note that
	// it's up to the caller to determine if the first largest
	// charge is the last item in the bank!
	// Returns the step index of the largest found
	largest := 0
	largestIdx := -1
	for i := start; i < end; i++ {
		if bb.Cells[i] > largest {
			largest = bb.Cells[i]
			largestIdx = i
		}
	}
	return largest, largestIdx
}

func (bb *BatteryBank) TotalCharge() int {
	// This will return the total charge of the battery bank
	// by using the logic provided by the challenge. This would
	// add up two numbers, largest charge then the second largest
	// RIGHT AFTER the finding of the largest charge.

	initialCharge, firstIdx := bb.LargestCharge(0, len(bb.Cells)-1)
	secondaryCharge, _ := bb.LargestCharge(firstIdx+1, len(bb.Cells))
	// fmt.Printf("Debug: Initial largest %d at index %d, secondary largest %d\n", initialCharge, firstIdx, secondaryCharge)

	// Now we have to concatenate the charges and return the integer
	concatenated := fmt.Sprintf("%d%d", initialCharge, secondaryCharge)
	result, _ := strconv.Atoi(concatenated)
	return result
}

func (bb *BatteryBank) FindTwelveHighest() []int {
	// This function will find the indexes of the 12 highest integers
	// in the battery bank and return them as a slice of integers. The
	// integers will be the indeces of the numbers found. I'll sort them
	// when I return them.
	// this will be the return value pre-sorted
	highestIdxs := make([]int, 0)

	// Keep track of used indeces
	usedIdxs := make(map[int]bool)

	for len(highestIdxs) < 12 {
		largest := 0
		largestIdx := -1
		for i, val := range bb.Cells {
			// Skip used indeces
			if usedIdxs[i] {
				continue
			}
			if val > largest {
				largest = val
				largestIdx = i
			}
		}
		if largestIdx == -1 {
			break
		}
		highestIdxs = append(highestIdxs, largestIdx)
		usedIdxs[largestIdx] = true
	}

	// Now to sort the indeces
	slices.Sort(highestIdxs)
	return highestIdxs
}

func (bb *BatteryBank) TotalChargePartTwo() int {
	// This will use the logic for part two of the challenge to
	// hopefully calculate the proper total charge. Let's see!
	highestIdxs := bb.FindTwelveHighest()

	// Now to concatenate the values found at those indeces
	concatenated := ""
	for _, idx := range highestIdxs {
		concatenated += fmt.Sprintf("%d", bb.Cells[idx])
	}
	fmt.Printf("Debug: Part Two concatenated string: %s\n", concatenated)
	result, _ := strconv.Atoi(concatenated)
	return result
}

// I found a better way. I'll use a stack to find the 12 highest values.
// I just so happen to have coded a stack already, and it's in the utils package
func (bb *BatteryBank) GetHighestPossible(numberLength int) int {
	// challenge calls for 12, but I won't hard code it in the function
	s := utils.NewGStack[int]()

	// Note that we'll check every single number, and only pop if the number is
	// less than the number at the top of the stack AND if the stack length plus
	// the remaining numbers to check is greater than the desired length.
	for i, val := range bb.Cells {
		// While stack is not empty, and top of stack is less than current value
		// and the stack length + remaining numbers is greater than desired length
		remaining := len(bb.Cells) - i

		// Pop all smaller values from the stack if we have enough remaining
		for !s.IsEmpty() {
			top, _ := s.Peek()
			if top < val && (s.Size()+remaining > numberLength) {
				s.Pop()
			} else {
				break
			}
		}

		// Only push if we haven't reached the desired length
		if s.Size() < numberLength {
			s.Push(val)
		}
	}

	// fmt.Printf("Debug: Stack size after processing: %d\n", s.Size())
	concatenated := ""
	for !s.IsEmpty() {
		top, _ := s.Pop()
		concatenated = fmt.Sprintf("%d%s", top, concatenated)
	}

	// fmt.Printf("Debug: Part Two concatenated string: %s\n", concatenated)

	result, _ := strconv.Atoi(concatenated)
	return result
}
