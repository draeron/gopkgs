package color

import (
	"math"
)

type HSL struct {
	H    uint16
	S, L uint8
}

func (l HSL) IsSame(r HSL) bool {
	return l.H == r.H && l.L == r.L && l.S == r.S
}

func FromF1(h, s, l uint8) HSV {
	return HSV{
		H: uint16(math.Round(float64(h) / 128. * 360)),
		S: uint8(math.Round(float64(s) / 128. * 100)),
		V: uint8(math.Round(float64(l) / 128. * 100)),
	}
}

func (c HSL) Equal(col Color) bool {
	return c.RGB().Equal(col)
}

func (c HSL) HSL() HSL {
	return c
}

func (c HSL) HSV() HSV {
	// todo optimise this
	return c.RGB().HSV()
}

func (c HSL) RGBA() (r, g, b, a uint32) {
	return c.RGB().RGBA()
}

func (c HSL) RGB() RGB {

	h := float64(c.H) / 360.
	s := float64(c.S) / 100.
	l := float64(c.L) / 100.

	var r, g, b float64

	if c.S == 0 {
		val := uint8(math.Round(l * 255))
		return RGB{
			val, val, val, 255,
		}
	}

	var q, p float64

	if l < 0.5 {
		q = l * (1 + s)
	} else {
		q = l + s - l*s
	}
	p = 2*l - q

	r = hue2rgb(p, q, h+1/3.)
	g = hue2rgb(p, q, h)
	b = hue2rgb(p, q, h-1/3.)

	return RGB{
		R: uint8((math.Round(r * 255))),
		G: uint8((math.Round(g * 255))),
		B: uint8((math.Round(b * 255))),
		A: 255,
	}
}

func hue2rgb(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}

	if t < 1/6. {
		return p + (q-p)*6*t
	}
	if t < 1/2. {
		return q
	}
	if t < 2/3. {
		return p + (q-p)*(2/3.-t)*6.
	}
	return p
}
