package geo

import (
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/xy"
	"image"
	"math"
)

var TwoPi = math.Pi * 2

func Coord(x, y float64) geom.Coord {
	return geom.Coord{x, y}
}

func CoordX(x float64) geom.Coord {
	return geom.Coord{x, 0}
}

func CoordY(y float64) geom.Coord {
	return geom.Coord{0, y}
}

func boundToRect(bounds *geom.Bounds) image.Rectangle {
	return image.Rectangle{
		image.Point{int(bounds.Min(0)), int(bounds.Min(1))},
		image.Point{int(bounds.Max(0)), int(bounds.Max(1))},
	}
}

func Radians(degrees int) float64 {
	return float64(degrees) * math.Pi / 180
}

func RadiansF(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func Degrees(radians float64) float64 {
	return radians * 180 / math.Pi
}

func Normalize(rad float64) float64 {
	return xy.NormalizePositive(rad)
}
