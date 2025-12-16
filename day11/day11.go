package main

import (
	"day11/reactor"
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
	lines, err := utils.GetFileLines(*filePtr)
	if err != nil {
		fmt.Println("Fatal:", err)
	}

	// giant text blob:
	// challengeText, err := utils.GetTextBlob(*filePtr)
	// if err != nil {
	//     fmt.Println("Fatal:", err)
	// }

	// Insert code here

	devices := []*reactor.Device{}
	for _, line := range lines {
		device := reactor.NewDevice(line)
		devices = append(devices, device)
	}

	// For easier navigation, I'll create a map of device names
	deviceMap := map[string]*reactor.Device{}

	for _, device := range devices {
		deviceMap[device.Name] = device
	}

	visited := map[string]bool{}
	totalPaths := reactor.TraverseDFS(deviceMap["you"], "out", deviceMap, visited)
	fmt.Printf("Total paths from 'you' to 'out' for Part One: %d\n", totalPaths)

	totalPathsPartTwo := reactor.PartTwoTraverse(deviceMap, "svr", "out")
	fmt.Printf("Total paths from 'srv' to 'out' that pass through 'dac' and 'fft' for Part Two: %d\n", totalPathsPartTwo)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
