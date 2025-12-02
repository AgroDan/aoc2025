package ranges

import "fmt"

type Range struct {
	Start int
	End   int
}

func NewRange(rangeText string) (*Range, error) {
	// rangeText should be "x-y"
	var r Range
	// going to be a little pedantic
	n, err := fmt.Sscanf(rangeText, "%d-%d", &r.Start, &r.End)
	if err != nil || n != 2 {
		return &Range{}, fmt.Errorf("Invalid range text: %s", rangeText)
	}
	return &r, nil
}

func (r Range) GetAllInvalidIds() []int {
	retval := []int{}
	for i := r.Start; i <= r.End; i++ {
		invalids := FindMirroredNumbers(i)
		if invalids > 0 {
			retval = append(retval, invalids)
		}
	}
	return retval
}

func (r Range) GetAllInvalidIdsPart2() []int {
	retval := []int{}
	for i := r.Start; i <= r.End; i++ {
		invalids := FindRepeatedPatterns(i)
		if invalids > 0 {
			retval = append(retval, invalids)
		}
	}
	return retval
}
