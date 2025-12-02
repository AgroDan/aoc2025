package ranges

import "strings"

func ParseRanges(input string) []string {
	// This will parse the full ranges and then
	// return a slice of strings signifying the
	// ranges, that we'll use to create a struct

	// This is a little too pythonic for me but it's early so...
	return strings.Split(strings.TrimSpace(input), ",")
}
