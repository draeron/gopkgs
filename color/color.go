package color

import (
	"github.com/draeron/gopkgs/logger"
	"image/color"
)

//go:generate go-enum -f=$GOFILE --noprefix --names

// Color x ENUM(
// Red,
// Orange,
// Yellow,
// YellowGreen,
// Green,
// CyanGreen,
// Cyan,
// CyanBlue,
// Blue,
// Purple,
// Magenta,
// MagentaRed,
// White,
// Black,
// LightGray,
// Gray,
// DarkGray,
// Transparent,
// )
type PaletteColor int

var log = logger.New("color")

type Color interface {
	HSL() HSL
	HSV() HSV
	RGB() RGB
	RGBA() (r, g, b, a uint32)
	Equal(c Color) bool
}

func FromColor(c color.Color) Color {
	r, g, b, a := c.RGBA()
	// Convert 16 bits channel to 8 bits
	return RGB{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: uint8(a >> 8),
	}
}

func (c PaletteColor) Lerp(to Color, t float32) Color {
	return c.RGB().Lerp(FromColor(to).RGB(), t)
}

func (c PaletteColor) Equal(col Color) bool {
	return c.RGB().Equal(col)
}

func (c PaletteColor) RGBA() (r, g, b, a uint32) {
	return c.RGB().RGBA()
}

func (c PaletteColor) HSL() HSL {
	if int(c) >= len(Palette) {
		log.Warn("invalid color index: ", c)
		return Palette[Black].HSL()
	} else {
		return Palette[c].HSL()
	}
}

func (c PaletteColor) HSV() HSV {
	if int(c) >= len(Palette) {
		log.Warn("invalid color index: ", c)
		return Palette[Black].HSV()
	} else {
		return Palette[c].HSV()
	}
}

func (c PaletteColor) RGB() RGB {
	if int(c) >= len(Palette) {
		log.Warn("invalid color index: ", c)
		return Palette[Black].RGB()
	} else {
		return Palette[c].RGB()
	}
}

func Colors() []PaletteColor {
	var cols []PaletteColor
	for cl, _ := range _PaletteColorMap {
		cols = append(cols, cl)
	}
	return cols
}

func (rc RGB) IsSame(r RGB) bool {
	return rc.R == r.R && rc.G == r.G && rc.B == r.B
}
