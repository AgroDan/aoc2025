package dial

type Dial struct {
	pos int
}

func NewDial(pos int) *Dial {
	if pos < 0 || pos > 99 {
		panic("Dial position must be between 0 and 99 inclusive")
	}
	return &Dial{pos: pos}
}

func (d *Dial) GetPosition() int {
	return d.pos
}

func (d *Dial) TurnRight(steps int) {
	// Changes the dial position to the right by the number of steps
	// assigned to it. This does _not_ return the amount of times
	// this dial _passes_ 0 like I initially thought. That is left
	// as an instruction for the loop that calls this method.
	// I should probably use the Euclidian modulus here but
	// I'm lazy. Y'know what's weird, Copilot is even suggesting
	// things to write in my annotations. I'm not sure what's
	// stranger, the fact that it does it or the fact that it's
	// pretty spot on. Yeah Copilot I do feel pretty lazy, what of it?
	for i := 0; i < steps; i++ {
		d.pos++
		if d.pos > 99 {
			d.pos = 0
		}
	}
}

func (d *Dial) TurnLeft(steps int) {
	// Like the TurnRight method, but in the opposite direction
	for i := 0; i < steps; i++ {
		d.pos--
		if d.pos < 0 {
			d.pos = 99
		}
	}
}

// I just had these previously because I mis-read the challenge, but this
// is what I did before! This is the second half of the challenge, so I'm
// just going to create new functions instead of retro-fitting the last ones.

func (d *Dial) TurnRightCountingRotations(steps int) int {
	passes := 0
	for i := 0; i < steps; i++ {
		d.pos++
		if d.pos > 99 {
			d.pos = 0
			passes++
		}
	}
	return passes
}

func (d *Dial) TurnLeftCountingRotations(steps int) int {
	// slightly different but basically the same
	passes := 0
	for i := 0; i < steps; i++ {
		d.pos--
		if d.pos == 0 {
			passes++
		}
		if d.pos < 0 {
			d.pos = 99
		}
	}
	return passes
}
