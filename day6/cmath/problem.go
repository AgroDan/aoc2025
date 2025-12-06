package cmath

import (
	"fmt"
	"strconv"
	"utils"
)

// Ahhh cephalopod math. To make this easier for me to grasp,
// I'm going to create a struct consisting of the _data_ around
// a particular problem. Given the current data looking like this:

// 123 328  51 64
//  45 64  387 23
//   6 98  215 314
// *   +   *   +

// I can use the operand on the bottom line to denote the bottom-left
// most point in the problem. Then I can create a grid of this data
// as a struct, and build information from that however I'd like.

type Problem struct {
	Graph   [][]string
	Operand string
}

func (p Problem) Print() {
	for _, row := range p.Graph {
		for _, val := range row {
			fmt.Printf("%s", val)
		}
		fmt.Printf("\n")
	}
}

func (p Problem) Solve() int {
	result := 0

	// First, build the numbers, remember right-to-left, but top down!
	var nums = make([]int, 0)
	for col := len(p.Graph[0]) - 1; col >= 0; col-- {
		builtNum := ""
		for row := 0; row < len(p.Graph)-1; row++ {
			if p.Graph[row][col] == " " {
				continue
			}
			builtNum = fmt.Sprintf("%s%s", builtNum, p.Graph[row][col])
		}
		num, err := strconv.Atoi(builtNum)
		if err != nil {
			fmt.Println("Error converting to int:", err)
		}
		nums = append(nums, num)
	}

	// Now apply the operand
	switch p.Operand {
	case "+":
		for _, n := range nums {
			result += n
		}
	case "*":
		result = 1
		for _, n := range nums {
			result *= n
		}
	default:
		fmt.Println("Unknown operand:", p.Operand)
	}

	return result
}

// So given the left-bottom most point of the problem and the right top
// coordinates, I can just take all the data from it and build the struct.
// And the cool thing is that I really don't need to find the top coord,
// because it'll just be the top row. I just need to find the _next_ operator,
// OR the end of the row.

func GetProblems(r utils.Runemap) []Problem {
	// Now find all the problem sets. Start from the bottom row, and seek ahead
	// until we find the next operator OR the end of the line.
	var problems []Problem
	bottomY := r.Height() - 1
	x := 0
	for x < r.Width() {
		runeAtX, err := r.Get(utils.Coord{X: x, Y: bottomY})
		if err != nil {
			panic("Error getting rune at bottom row: " + err.Error())
		}
		if runeAtX == ' ' {
			// Not an operand, so keep looking
			x++
			continue
		}

		// otherwise, we found an operand AND the left-most column in this problem.
		// Now find the right-most.
		seekX := x + 1
		var rightMostX int
		for {
			if seekX >= r.Width() {
				// hit the end of the file, so this is the right-most column
				rightMostX = seekX - 1
				break
			}
			runeAtSeek, err := r.Get(utils.Coord{X: seekX, Y: bottomY})
			if err != nil {
				panic("Error getting rune at seek:" + err.Error())
			}
			if runeAtSeek != ' ' {
				// found the next operand, so the right-most column is one before this
				// also keep in mind there's a column of whitespace between problems,
				// hence -2
				rightMostX = seekX - 2
				break
			}
			seekX++
		}

		// Now we can build a problem with the left and right-most coords.
		var graph [][]string
		for row := 0; row <= bottomY; row++ {
			var rowData []string
			for col := x; col <= rightMostX; col++ {
				runeAtPos, err := r.Get(utils.Coord{X: col, Y: row})
				if err != nil {
					panic("Error getting rune at position:" + err.Error())
				}
				rowData = append(rowData, string(runeAtPos))
			}
			graph = append(graph, rowData)
		}

		// Get the operand at the bottom row
		operandRune, err := r.Get(utils.Coord{X: x, Y: bottomY})
		if err != nil {
			fmt.Println("Error getting operand rune:", err)
			break
		}
		operand := string(operandRune)
		newProblem := Problem{
			Graph:   graph,
			Operand: operand,
		}
		problems = append(problems, newProblem)

		// Now move x to the right-most + 2 (to skip the whitespace)
		x = rightMostX + 2
	}

	return problems
}
