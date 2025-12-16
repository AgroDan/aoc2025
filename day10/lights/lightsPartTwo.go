package lights

import (
	"fmt"
	"strconv"
	"strings"
)

// So based on how I'm going to solve part two, I don't
// need to use bitwise math here. Instead this will just
// be integers. Since I parsed everything as bitmasks
// already, I'm going to basically re-parse into integers.
// Admittedly, this should be a lot easier to parse.

type JoltMachine struct {
	// note -- part 2 doesn't care about the bitwise lights
	Buttons [][]int
	Joltage []int
}

func NewJoltMachine(line string) *JoltMachine {
	m := &JoltMachine{}

	mSplit := strings.Split(line, " ")

	// buttons
	for i := 1; i < len(mSplit)-1; i++ {
		buttonStr := strings.Trim(mSplit[i], "()")
		buttonInd := strings.Split(buttonStr, ",")
		button := make([]int, len(buttonInd))
		for j, s := range buttonInd {
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			button[j] = num
		}
		m.Buttons = append(m.Buttons, button)
	}

	// joltage
	joltageStr := strings.Trim(mSplit[len(mSplit)-1], "{}")
	joltageInd := strings.Split(joltageStr, ",")
	for _, s := range joltageInd {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		m.Joltage = append(m.Joltage, num)
	}

	return m
}

func (m *JoltMachine) Print() {
	fmt.Printf("Buttons:\n")
	for i, b := range m.Buttons {
		fmt.Printf("  %d: %v\n", i, b)
	}
	fmt.Printf("Joltage: %v\n", m.Joltage)
}
