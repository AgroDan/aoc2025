package junctions

import "fmt"

// This will parse each object, given a string, and separate the
// values into 3 integers

func ParseJunction(input string) (*Junction, error) {
	var x, y, z int
	_, err := fmt.Sscanf(input, "%d,%d,%d", &x, &y, &z)
	if err != nil {
		return nil, err
	}
	return NewJunction(x, y, z), nil
}

func JunctionList(lines []string) ([]*Junction, error) {
	junctions := make([]*Junction, 0, len(lines))
	for _, line := range lines {
		junction, err := ParseJunction(line)
		if err != nil {
			return nil, err
		}
		junctions = append(junctions, junction)
	}
	return junctions, nil
}
