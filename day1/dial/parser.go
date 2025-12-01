package dial

import "strconv"

func (d *Dial) ParseInstruction(inst string) {
	// This will parse the direction and execute
	rotations, err := strconv.Atoi(inst[1:])
	if err != nil {
		panic("Something really weird happened and we shouldn't be here, figure it out nerd")
	}

	switch inst[0] {
	case 'R':
		d.TurnRight(rotations)
	case 'L':
		d.TurnLeft(rotations)
	default:
		panic("I couldn't read the instructions properly, GIT GUD NERD")
	}
}

func (d *Dial) ParseInstructionCountRotations(inst string) int {
	// This will parse the direction and how many steps we take,
	// and return the number of times we pass position 0
	rotations, err := strconv.Atoi(inst[1:])
	if err != nil {
		panic("Something really weird happened and we shouldn't be here, figure it out nerd")
	}

	fullRotations := 0
	switch inst[0] {
	case 'R':
		fullRotations += d.TurnRightCountingRotations(rotations)
	case 'L':
		fullRotations += d.TurnLeftCountingRotations(rotations)
	default:
		panic("I couldn't read the instructions properly, GIT GUD NERD")
	}
	return fullRotations
}
