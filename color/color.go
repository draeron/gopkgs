package color

import (
	"image/color"
	"sort"
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

type Color interface {
	HSL() HSL
	HSV() HSV
	RGB() RGB
	RGBA() (r, g, b, a uint32)
	Equal(c Color) bool
}

func FromStdColor(c color.Color) Color {
	r, g, b, a := c.RGBA()
	// Convert 16 bits channel to 8 bits
	return RGB{
		R: uint8(r >> 1),
		G: uint8(g >> 1),
		B: uint8(b >> 1),
		A: uint8(a >> 1),
	}
}

func FromInt32(c int32) color.Color {
	return RGB{
		R: uint8(c >> 24),
		G: uint8(c >> 16),
		B: uint8(c >> 8),
		A: uint8(c),
	}
}

func ToInt32(c color.Color) int32 {
	r, g, b, a := c.RGBA()
	// Convert 16 bits channel to 8 bits then pack into int32
	return int32(r>>8)<<24 | int32(g>>8)<<16 | int32(b>>8)<<8 | int32(a>>8)
}

func ToRBGA(c Color) color.RGBA {
	return color.RGBA(c.RGB())
}

func (c PaletteColor) Lerp(to Color, t float32) Color {
	return c.RGB().Lerp(FromStdColor(to).RGB(), t)
}

func (c PaletteColor) Equal(col Color) bool {
	return c.RGB().Equal(col)
}

func (c PaletteColor) RGBA() (r, g, b, a uint32) {
	return c.RGB().RGBA()
}

func (c PaletteColor) HSL() HSL {
	if int(c) >= len(Palette) {
		log.Warnf("invalid color index: ", c)
		return Palette[Black].HSL()
	} else {
		return Palette[c].HSL()
	}
}

func (c PaletteColor) HSV() HSV {
	if int(c) >= len(Palette) {
		log.Warnf("invalid color index: ", c)
		return Palette[Black].HSV()
	} else {
		return Palette[c].HSV()
	}
}

func (c PaletteColor) RGB() RGB {
	if int(c) >= len(Palette) {
		log.Warnf("invalid color index: ", c)
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
	sort.Slice(cols, func(i, j int) bool {
		return cols[i] < cols[j]
	})
	return cols
}

func (rc RGB) IsSame(r RGB) bool {
	return rc.R == r.R && rc.G == r.G && rc.B == r.B
}
