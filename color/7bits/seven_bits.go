package seven_bits

import (
	"github.com/draeron/gopkgs/color"
)

// color in 7 bits
type SevenColor struct {
	R, G, B uint8
}

func (bc SevenColor) HSL() color.HSL {
	return bc.RGB().HSL()
}

func (bc SevenColor) HSV() color.HSV {
	return bc.RGB().HSV()
}

func (bc SevenColor) RGB() color.RGB {
	return color.FromStdColor(&bc).RGB()
}

func (bc SevenColor) Equal(c color.Color) bool {
	return c.RGB().Equal(c)
}

func FromColor(color color.Color) SevenColor {
	r, g, b, _ := color.RGBA()
	return SevenColor{
		R: uint8(r >> 9),
		G: uint8(g >> 9),
		B: uint8(b >> 9),
	}
}

/*
	Convert 7 to 16 bits channels
*/
func (bc SevenColor) RGBA() (uint32, uint32, uint32, uint32) {
	return uint32(bc.R) << 9, uint32(bc.G) << 9, uint32(bc.B) << 9, 0xffff
}

func (bc SevenColor) IsSame(other SevenColor) bool {
	return bc.R == other.R && bc.G == other.G && bc.B == other.B
}
