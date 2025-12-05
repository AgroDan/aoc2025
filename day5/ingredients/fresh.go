package ingredients

import "strconv"

type IngredientRange struct {
	lowerBound, upperBound int
}

func NewIngredientRange(lower, upper int) IngredientRange {
	return IngredientRange{
		lowerBound: lower,
		upperBound: upper,
	}
}

func (ing IngredientRange) String() string {
	lb := strconv.Itoa(ing.lowerBound)
	ub := strconv.Itoa(ing.upperBound)
	return "[" + lb + "-" + ub + "]"
}

func (ing IngredientRange) IsFresh(value int) bool {
	// This just returns whether or not the provided
	// ingredient ID is considered fresh based on the given
	// ID and whether or not it falls within the provided range
	return value >= ing.lowerBound && value <= ing.upperBound
}

// to count the total amount of fresh IDs, I'll have to create
// a finite set so I can just count every single ingredient
// regardless if it's in an overlapping range.

func (ing IngredientRange) FindAllFreshIDs(s *set) {
	for i := ing.lowerBound; i <= ing.upperBound; i++ {
		s.Add(i)
	}
}
