package lights

import (
	"fmt"
	"strconv"
	"strings"
)

type Machine struct {
	Lights   uint64
	Desired  uint64
	lightlen int
	Buttons  []uint64
	Joltage  []int // not sure what this is
}

func NewMachine(line string) *Machine {
	m := &Machine{}

	splitStr := strings.Split(line, " ")

	// First, desired state
	desiredStr := strings.Trim(splitStr[0], "[]")
	m.lightlen = len(desiredStr)
	desiredMask := uint64(0)
	for _, s := range desiredStr {
		if s == '#' {
			desiredMask = (desiredMask << 1) | 1
		} else {
			desiredMask = (desiredMask << 1)
		}
	}
	m.Desired = desiredMask

	// reset the machine to zero to start
	m.Lights = 0

	// Now buttons

	for i := 1; i < len(splitStr)-1; i++ {
		buttonStr := strings.Trim(splitStr[i], "()")
		buttonMask := uint64(0)
		buttonInd := strings.Split(buttonStr, ",")
		for _, s := range buttonInd {
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			// Convert from left-indexed to right-indexed (bit position)
			bitPos := m.lightlen - 1 - num
			buttonMask |= (1 << uint(bitPos))
		}
		m.Buttons = append(m.Buttons, buttonMask)
	}

	// Finally, joltage (not sure what this is for)
	joltageStr := strings.Trim(splitStr[len(splitStr)-1], "{}")
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

func (m *Machine) Print() {
	fmt.Printf("Lights:  [%0*b]\n", m.lightlen, m.Lights)
	fmt.Printf("Desired: [%0*b]\n", m.lightlen, m.Desired)
	fmt.Printf("Buttons:\n")
	for i, b := range m.Buttons {
		fmt.Printf("  %d: [%0*b]\n", i, len(m.Buttons), b)
	}
	fmt.Printf("Joltage: %v\n", m.Joltage)
}
