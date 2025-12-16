package main

import (
	"day12/presents"
	"flag"
	"fmt"
	"time"
	"utils"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")
	// any additional flags add here

	flag.Parse()

	// Choose based on the challenge...
	// individual lines:
	// lines, err := utils.GetFileLines(*filePtr)
	// if err != nil {
	//     fmt.Println("Fatal:", err)
	// }

	// giant text blob:
	challengeText, err := utils.GetTextBlob(*filePtr)
	if err != nil {
		fmt.Println("Fatal:", err)
	}

	// Insert code here
	presents, regions := presents.ParseChallenge(challengeText)

	// for _, v := range presents {
	// 	v.Print()
	// 	fmt.Printf("------\n")
	// }

	// for _, r := range regions {
	// 	r.Print()
	// 	fmt.Printf("------\n")
	// }

	// fmt.Println("Rotating present at index 1 right")
	// p := presents[1].ClonePresent()
	// p.RotateRT()
	// p.Print()

	// fmt.Println("Rotating present at index 3 left")
	// p2 := presents[3].ClonePresent()
	// p2.RotateLT()
	// p2.Print()

	total := 0
	for _, r := range regions {
		// Let's just check to see if there's enough volume to fill the region
		target := r.Area()
		volSlice := make([]int, 0)

		for i, v := range r.PresentIdx {
			// Get the total volume of presents
			presents[i].Volume()
			// fmt.Printf("Present %d has volume %d, with a total volume of %d\n", i, presents[i].Volume(), presents[i].Volume()*v)
			volSlice = append(volSlice, presents[i].Volume()*v)
		}

		// i highly doubt this will work but let's see
		if target >= utils.SumIntSlice(volSlice) {
			// fmt.Printf("Region %v can be filled with total volume %d\n", r, utils.SumIntSlice(volSlice))
			total++
		}
	}

	fmt.Printf("Total regions that can be filled: %d\n", total)
	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
