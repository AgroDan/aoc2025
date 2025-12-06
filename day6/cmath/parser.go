package cmath

import (
	"strings"
	"utils"
)

// This is the parser to read all the data given in the challenge.
// Since formatting of this particular file is key here, I'll have
// to figure out a way to ensure that everything lines up I guess.
//
// Additionally, the datatype here really has to be a 2-dimensional
// slice of _strings_, since the types are both numbers AND operands.
// Yikes.

func ReadHomework(input []string) [][]string {
	var homework [][]string
	for _, line := range input {
		rowData := strings.Fields(line)
		homework = append(homework, rowData)
	}
	return homework
}

func PlotHomework(input []string) utils.Runemap {
	return utils.NewRunemap(input)
} // this is silly but STFU
