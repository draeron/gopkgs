package geo

import (
	"fyne.io/fyne"
	"github.com/twpayne/go-geom"
	"image"
)

type Rectangle struct {
	Point
	W, H float64
}

func NewRectangle(x, y, w, h float64) Rectangle {
	return Rectangle{Point{x, y}, w, h}
}

func (r Rectangle) End() Point {
	return Point{r.X + r.W, r.Y + r.W}
}

func (p *Rectangle) Intersect(ge Geometry) bool {
	return intersect(p, ge)
}

func (r Rectangle) ToImg() image.Rectangle {
	return image.Rect(
		int(r.X),
		int(r.Y),
		int(r.X+r.W),
		int(r.Y+r.H),
	)
}

func (r Rectangle) ToFyneSize() fyne.Size {
	return fyne.NewSize(int(r.W), int(r.H))
}

func FromImgRect(r image.Rectangle) Rectangle {
	return Rectangle{
		Point{float64(r.Min.X), float64(r.Min.Y)},
		float64(r.Max.X - r.Min.X),
		float64(r.Max.Y - r.Min.Y),
	}
}

func ToRect(bounds *geom.Bounds) Rectangle {
	return Rectangle{
		Point{bounds.Min(0), bounds.Min(1)},
		bounds.Max(0) - bounds.Min(0),
		bounds.Max(1) - bounds.Min(1),
	}
}
