package color

import (
	"image/color"
	"math"
)

type RGB color.RGBA

func (rc RGB) RGB() RGB {
	return rc
}

func (rc RGB) Lerp(to Color, t float32) RGB {
	target := FromColor(to).RGB()
	rc.R += uint8(float32(target.R-rc.R) * t)
	rc.G += uint8(float32(target.G-rc.G) * t)
	rc.B += uint8(float32(target.B-rc.B) * t)
	rc.A += uint8(float32(target.A-rc.A) * t)
	return rc
}

func (rc RGB) HSV() HSV {

	red := float64(rc.R)
	green := float64(rc.G)
	blue := float64(rc.B)

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

func (rc RGB) Lightness() uint8 {
	max := uint8(0)
	if rc.R > max {
		max = rc.R
	}
	if rc.G > max {
		max = rc.G
	}
	if rc.B > max {
		max = rc.B
	}
	return max
}

func (rc RGB) RGBA() (r, g, b, a uint32) {
	r = uint32(rc.R)
	r |= r << 8
	g = uint32(rc.G)
	g |= g << 8
	b = uint32(rc.B)
	b |= b << 8
	a = uint32(rc.A)
	a |= a << 8
	return
}

//func (c RGB) RGBA() (r, g, b, a uint32) {
//	return uint32(c.R) * 0xffff / 255,
//		uint32(c.G) * 0xffff / 255,
//		uint32(c.B) * 0xffff / 255,
//		0xffff
//}

func (rc RGB) HSL() HSL {

	r := rc.R
	g := rc.G
	b := rc.B

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

func (rc RGB) Equal(r Color) bool {
	c2 := r.RGB()
	return rc.R*rc.A == c2.R*c2.A && rc.G*rc.A == c2.G*c2.A && rc.B*rc.A == c2.B*c2.A
}
