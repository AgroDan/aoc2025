package presents

import (
	"strings"
)

func ParseChallenge(inData string) (map[int]Present, []Region) {
	// we'll split on double-line feeds
	data := strings.Split(inData, "\n\n")

	presents := make(map[int]Present)
	for i := 0; i < len(data)-1; i++ {
		lines := strings.Split(data[i], "\n")
		present := ParsePresent(lines)

		presents[present.index] = present
		// fmt.Printf("Parsed present index %d with shape: %v\n", present.index, present.shape)
	}

	challenges := make([]Region, 0)

	challengeLines := strings.Split(data[len(data)-1], "\n")
	for _, line := range challengeLines {
		region := ParseRegion(line)
		challenges = append(challenges, region)
		// fmt.Printf("Parsed region height %d width %d with presents %v\n", region.height, region.width, region.presentIdx)
	}
	return presents, challenges
}
