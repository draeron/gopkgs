package color

import (
	"math"
)

type HSV struct {
	H    uint16
	S, V uint8
}

func (l HSV) HSV() HSV {
	return l
}

func (l HSV) IsSame(r HSV) bool {
	return l.H == r.H && l.V == r.V && l.S == r.S
}

func (c HSV) Equal(col Color) bool {
	return c.RGB().Equal(col)
}

func (c HSV) RGBA() (r, g, b, a uint32) {
	return c.RGB().RGBA()
}

func (c HSV) HSL() HSL {
	// todo optimise this
	return c.RGB().HSL()
}

func (c HSV) RGB() RGB {
	var red, green, blue float64

	h := float64(c.H)
	s := float64(c.S) / 100.
	b := float64(c.V) / 100 * 255.

	if c.S == 0 {
		val := uint8(b)
		return RGB{
			val, val, val, 255,
		}
	}

	var sextant = math.Floor(h / 60)
	var remainder = (h / 60) - sextant
	var val1 = b * (1 - s)
	var val2 = b * (1 - (s * remainder))
	var val3 = b * (1 - (s * (1 - remainder)))

	switch sextant {
	case 1:
		red = val2
		green = b
		blue = val1

	case 2:
		red = val1
		green = b
		blue = val3

	case 3:
		red = val1
		green = val2
		blue = b

	case 4:
		red = val3
		green = val1
		blue = b

	case 5:
		red = b
		green = val1
		blue = val2

	case 6, 0:
		red = b
		green = val3
		blue = val1
	}

	return RGB{
		R: uint8((math.Round(red))),
		G: uint8((math.Round(green))),
		B: uint8((math.Round(blue))),
		A: 255,
	}
}
