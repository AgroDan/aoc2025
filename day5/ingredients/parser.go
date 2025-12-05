package ingredients

import (
	"strconv"
	"strings"
)

func ParseIngredients(input string) ([]IngredientRange, []int) {
	// Returns a slice of IngredientRanges, as well as the
	// list of ingredient IDs.

	// for return values
	var ingredientRanges []IngredientRange
	var ingredientIDs []int

	// first, biforcate the input based on newlines
	splitInput := strings.Split(input, "\n\n")
	rangeLines := strings.Split(splitInput[0], "\n")
	for _, line := range rangeLines {
		r := strings.Split(line, "-")
		left, _ := strconv.Atoi(r[0])
		right, _ := strconv.Atoi(r[1])
		ingredientRanges = append(ingredientRanges, NewIngredientRange(left, right))
	}

	idLines := strings.Split(splitInput[1], "\n")
	for _, line := range idLines {
		id, _ := strconv.Atoi(line)
		ingredientIDs = append(ingredientIDs, id)
	}

	return ingredientRanges, ingredientIDs
}
