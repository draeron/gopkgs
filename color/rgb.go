package color

import (
	"image/color"
	"math"
)

type RGB color.RGBA

func (l RGB) RGB() RGB {
	return l
}

func (rgb RGB) HSV() HSV {

	red := float64(rgb.R)
	green := float64(rgb.G)
	blue := float64(rgb.B)

	var max = math.Max(math.Max(red, green), blue)
	var min = math.Min(math.Min(red, green), blue)

	var hue, saturation, value float64
	value = max

	if min == max {
		hue = 0
		saturation = 0
	} else {
		var delta = max - min
		saturation = delta / max

		if red == max {
			hue = (green - blue) / delta
		} else if green == max {
			hue = 2 + ((blue - red) / delta)
		} else {
			hue = 4 + ((red - green) / delta)
		}
		hue *= 60
		if hue < 0 {
			hue += 360
		}
		if hue > 360 {
			hue -= 360
		}
	}

	return HSV{
		uint16(hue), uint8(saturation * 100), uint8(value),
	}
}

func (c RGB) Lightness() uint8 {
	max := uint8(0)
	if c.R > max {
		max = c.R
	}
	if c.G > max {
		max = c.G
	}
	if c.B > max {
		max = c.B
	}
	return max
}

func (c RGB) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = uint32(c.A)
	a |= a << 8
	return
}

//func (c RGB) RGBA() (r, g, b, a uint32) {
//	return uint32(c.R) * 0xffff / 255,
//		uint32(c.G) * 0xffff / 255,
//		uint32(c.B) * 0xffff / 255,
//		0xffff
//}

func (c RGB) HSL() HSL {

	r := c.R
	g := c.G
	b := c.B

	var fH, fS, fL float64

	fR := float64(r) / 255
	fG := float64(g) / 255
	fB := float64(b) / 255
	max := math.Max(math.Max(fR, fG), fB)
	min := math.Min(math.Min(fR, fG), fB)
	fL = (max + min) / 2
	if max == min {
		// Achromatic.
		fH, fS = 0, 0
	} else {
		// Chromatic.
		d := max - min
		if fL > 0.5 {
			fS = d / (2.0 - max - min)
		} else {
			fS = d / (max + min)
		}
		switch max {
		case fR:
			fH = (fG - fB) / d
			if fG < fB {
				fH += 6
			}
		case fG:
			fH = (fB-fR)/d + 2
		case fB:
			fH = (fR-fG)/d + 4
		}
		fH /= 6
	}
	return HSL{
		H: uint16(math.Round(fH * 360)),
		S: uint8(math.Round(fS * 100)),
		L: uint8(math.Round(fL * 100)),
	}
}

func (c RGB) Equal(r Color) bool {
	c2 := r.RGB()
	return c.R*c.A == c2.R*c2.A && c.G*c.A == c2.G*c2.A && c.B*c.A == c2.B*c2.A
}
