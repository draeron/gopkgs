package geo

import (
	"github.com/draeron/gopkg/color"
	"github.com/fogleman/gg"
	"github.com/twpayne/go-geom"
)

type Arc struct {
	Circle
	Angle float64
	Width float64
}

func NewArc() Arc {
	return Arc{
		Circle{Zero, 0},
		0, TwoPi,
	}
}

func (a Arc) Pos() Point {
	return a.Center
}

func (a Arc) Centroid() Point {
	return a.Center
}

func (a Arc) Intersect(geom Geometry) bool {
	return intersect(&a, geom)
}

func (a *Arc) Translate(p Point) {
	a.Circle.Translate(p)
}

func (a *Arc) Rotate(rad float64) {
	a.RotateAround(rad, a.Center)
}

func (a *Arc) RotateAround(rad float64, anchor Point) {
	a.Circle.RotateAround(rad, anchor)
	a.Angle = Normalize(a.Angle + rad)
}

func (a Arc) Bounds() Rectangle {
	return a.Circle.Bounds()
}

func (a Arc) Draw(g *gg.Context) {
	g.Push()
	a.DrawOutline(g)
	a.DrawCenter(g)
	g.Pop()
}

func (a Arc) DrawOutline(g *gg.Context) {
	g.Push()

	g.DrawArc(a.Center.X, a.Center.Y, a.Radius, a.Angle, a.Angle+a.Width)
	g.Stroke()

	g.SetColor(color.Green)
	pts := a.Center.AddXY(a.Radius, 0)
	pts.RotateAround(a.Angle, a.Center)
	pts.Draw(g)

	g.SetColor(color.Red)
	pts = a.Center.AddXY(a.Radius, 0)
	pts.RotateAround(a.Angle+a.Width, a.Center)
	pts.Draw(g)

	g.Pop()
}

func (a Arc) DrawCenter(g *gg.Context) {
	a.Circle.DrawCenter(g)
}

func (a Arc) bounds() *geom.Bounds {
	return a.Circle.bounds()
}
