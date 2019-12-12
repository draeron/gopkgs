package color

import (
	"github.com/draeron/gopkg/logger"
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
	return RGB{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}
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

func (l RGB) IsSame(r RGB) bool {
	return l.R == r.R && l.G == r.G && l.B == r.B
}
