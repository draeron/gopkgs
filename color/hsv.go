package color

import (
	"math"
)

type HSV struct {
	H    uint16
	S, V uint8
}

func (hv HSV) HSV() HSV {
	return hv
}

func (hv HSV) Lerp(to HSV, t float32) HSV {
	// Hue interpolation
	h := float32(0)
	d := float32(to.H - hv.H)

	if hv.H > to.H {
		hv.H, to.H = to.H, hv.H // Swap
		d = -d
		t = 1 - t
	}

	if d > 180 { // 180deg
		hv.H = 360 - hv.H // 360deg
		h = (float32(hv.H) + t * float32(to.H - hv.H)) // 360deg
		h = float32(uint16(h) % 360)
	}
	if d <= 180 { // 180deg
		h = float32(hv.H) + t * d
	}

	return HSV{
		H: uint16(h),
		S: uint8(float32(hv.S) + t * float32(to.S - hv.S)),
		V: uint8(float32(hv.V) + t * float32(to.V - hv.V)),
	}
}

func (hv HSV) IsSame(r HSV) bool {
	return hv.H == r.H && hv.V == r.V && hv.S == r.S
}

func (hv HSV) Equal(col Color) bool {
	return hv.RGB().Equal(col)
}

func (hv HSV) RGBA() (r, g, b, a uint32) {
	return hv.RGB().RGBA()
}

func (hv HSV) HSL() HSL {
	// todo optimise this
	return hv.RGB().HSL()
}

func (hv HSV) RGB() RGB {
	var red, green, blue float64

	h := float64(hv.H)
	s := float64(hv.S) / 100.
	b := float64(hv.V) / 100 * 255.

	if hv.S == 0 {
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
