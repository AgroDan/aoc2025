package ingredients

import (
	"cmp"
	"slices"
)

// This is a job for the sweeping line algorithm!

func SweepIngredientRanges(ranges []IngredientRange) int {
	// This function will just loop through all the ingredient ranges
	// and add up all the differences between the upper and lower bounds
	// by way of the sweeping line algorithm, something I just discovered!
	type event struct {
		position, typ int // typ: 1 for start, -1 for end
	}
	var events []event
	for _, r := range ranges {
		events = append(events, event{position: r.lowerBound, typ: 1})
		events = append(events, event{position: r.upperBound + 1, typ: -1}) // +1 to make end exclusive
	}

	if len(events) == 0 {
		return 0
	}

	// Sort events by position, then by type
	slices.SortFunc(events, func(a, b event) int {
		if a.position == b.position {
			return cmp.Compare(a.typ, b.typ)
		}
		return cmp.Compare(a.position, b.position)
	})

	activeRanges := 0
	prevPosition := events[0].position
	totalFresh := 0

	for _, e := range events {
		if activeRanges > 0 {
			totalFresh += e.position - prevPosition
		}
		activeRanges += e.typ
		prevPosition = e.position
	}

	return totalFresh

}
