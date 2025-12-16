package lights

import "fmt"

// func (m *Machine) stateBitmask() uint16 {
// 	var bitmask uint16 = 0
// 	for i, light := range m.Lights {
// 		if light {
// 			bitmask |= (1 << i)
// 		}
// 	}
// 	return bitmask
// }

// func (m *Machine) desiredBitmask() uint16 {
// 	var bitmask uint16 = 0
// 	for i, light := range m.Desired {
// 		if light {
// 			bitmask |= (1 << i)
// 		}
// 	}
// 	return bitmask
// }

func (m *Machine) Solve() (int, error) {
	start := m.Lights
	target := m.Desired

	delta := start ^ target

	// Initialize no solution state
	bestCount := 99999999999999
	ok := false

	// Total combos of buttons are 2^n
	n := len(m.Buttons)
	totalCombos := uint64(1) << n

	for mask := uint64(0); mask < totalCombos; mask++ {
		var current uint64 = 0

		// Apply all button effects included with mask
		for i := 0; i < n; i++ {
			if (mask & (1 << uint(i))) != 0 {
				current ^= m.Buttons[i]
			}
		}

		// Check if this is a solution
		if current == delta {
			ok = true
			// Count bits in mask
			count := 0
			for i := 0; i < n; i++ {
				if (mask & (1 << uint(i))) != 0 {
					count++
				}
			}
			if count < bestCount {
				bestCount = count
			}
		}
	}
	if !ok {
		return -1, fmt.Errorf("no solution found")
	}

	return bestCount, nil
}
