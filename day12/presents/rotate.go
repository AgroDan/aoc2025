package presents

// Not sure if I'll need this or not...but just in case.

func (p Present) ClonePresent() Present {
	var newShape [3][3]rune
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			newShape[i][j] = p.shape[i][j]
		}
	}
	return Present{
		index: p.index,
		shape: newShape,
	}
}

func (p *Present) RotateRT() {
	// This just rotates the present clockwise 90 degrees.
	var newShape [3][3]rune
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			newShape[x][2-y] = p.shape[y][x]
		}
	}
	p.shape = newShape
}

func (p *Present) RotateLT() {
	// rotates counter-clockwise, changed to LT and RT because CW anc CCW were too similar
	var newShape [3][3]rune
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			newShape[2-x][y] = p.shape[y][x]
		}
	}
	p.shape = newShape
}

func (p *Present) FlipVertical() {
	// flips the present vertically
	var newShape [3][3]rune
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			newShape[2-y][x] = p.shape[y][x]
		}
	}
	p.shape = newShape
}

func (p *Present) FlipHorizontal() {
	// flips the present horizontally
	var newShape [3][3]rune
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			newShape[y][2-x] = p.shape[y][x]
		}
	}
	p.shape = newShape
}
