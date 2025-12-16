package reactor

// This is going to be my attempt at using Dynamic Programming, or
// dynamic digital programming, which is supposedly faster at counting
// gigantic numbers of combinations, so let's see I guess.

const (
	DAC = 0b0001 // if we see a DAC, set LSB
	FFT = 0b0010 // if we see FFT, second bit flagged
)

func seenBit(node string) int64 {
	// this will just flip the appropriate bits if we spot a DAC or FFT
	mask := int64(0)
	if node == "dac" {
		mask |= DAC
	}

	if node == "fft" {
		mask |= FFT
	}
	return mask
}

func dpTraverse(deviceMap *map[string]*Device, current string, target string, memo map[string]map[int64]int, mask int64, visited map[string]map[int64]bool) int {
	// if we reached the target, return ONLY if we've seen both DAC and FFT
	if current == target {
		if mask == (DAC | FFT) {
			return 1
		}
		return 0
	}

	// memo check!
	if _, ok := memo[current]; ok {
		if v, ok2 := memo[current][mask]; ok2 {
			return v
		}
	} else {
		memo[current] = make(map[int64]int)
	}

	if _, ok := visited[current]; !ok {
		visited[current] = make(map[int64]bool)
	}

	// this will be the ultimate test to determine if this path is cyclical.
	// it shouldn't be, but if it is, this will let me know.
	if visited[current][mask] {
		panic("Cycle detected in DP traversal!")
	}
	visited[current][mask] = true

	var totalPaths int = 0

	for _, connName := range (*deviceMap)[current].Connections {
		newMask := mask | seenBit(connName)
		totalPaths += dpTraverse(deviceMap, connName, target, memo, newMask, visited)
	}

	// unmark visited for other paths
	visited[current][mask] = false

	// store in memo
	memo[current][mask] = totalPaths
	return totalPaths
}

func PartTwoTraverse(deviceMap map[string]*Device, start string, target string) int {
	memo := make(map[string]map[int64]int)
	visited := make(map[string]map[int64]bool)
	return dpTraverse(&deviceMap, start, target, memo, 0, visited)
}
