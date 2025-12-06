package cmath

import (
	"fmt"
	"strconv"
)

// Here is where we actually do the homework.

func DoHomework(homework [][]string) int {
	total := 0

	// Going to loop over every single column...so in reality
	// I just need to loop over the first _row_ per se, and just
	// loop downwards

	// Excuse my wording, this is for my own clarity
	for col := 0; col < len(homework[0]); col++ {
		result := 0
		switch homework[len(homework)-1][col] {
		case "+":
			for row := 0; row < len(homework)-1; row++ {
				num, err := strconv.Atoi(homework[row][col])
				if err != nil {
					fmt.Println("Error converting to int:", err)
				}
				result += num
			}
		case "*":
			result = 1
			for row := 0; row < len(homework)-1; row++ {
				num, err := strconv.Atoi(homework[row][col])
				if err != nil {
					fmt.Println("Error converting to int:", err)
				}
				result *= num
			}
		default:
			fmt.Println("Unknown operator:", homework[len(homework)-1][col])
		}
		total += result
	}

	// // we're going to loop through each COLUMN, which is
	// // the second index in the homework variable.
	// // homework[row][col]c
	// for row := 0; row < len(homework); row++ {
	// 	result := 0
	// 	switch homework[row][len(homework[row])-1] {
	// 	case "+":
	// 		// we're adding here
	// 		for col := 0; col < len(homework[row])-2; col++ {
	// 			num, err := strconv.Atoi(homework[row][col])
	// 			if err != nil {
	// 				fmt.Println("Error converting to int:", err)
	// 			}
	// 			result += num
	// 		}
	// 	case "*":
	// 		// we're multiplying now, so start with 1 so it doesn't
	// 		// zero out the end result
	// 		result = 1
	// 		for col := 0; col < len(homework[row])-2; col++ {
	// 			num, err := strconv.Atoi(homework[row][col])
	// 			if err != nil {
	// 				fmt.Println("Error converting to int:", err)
	// 			}
	// 			result *= num
	// 		}
	// 	default:
	// 		fmt.Println("Unknown operator:", homework[row][len(homework[row])-1])
	// 	}
	// 	total += result
	// }
	return total
}
