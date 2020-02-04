package geo

import (
	"github.com/draeron/gopkgs/color"
	"github.com/fogleman/gg"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/xy"
)

type Line struct {
	Start, End Point
}

func NewLine(start, end Point) Line {
	return Line{start, end}
}

func (l Line) Intersect(ge Geometry) bool {
	return intersect(&l, ge)
}

func (l Line) Bounds() Rectangle {
	return ToRect(l.bounds())
}

func (l Line) bounds() *geom.Bounds {
	return l.toGeom().Bounds()
}

func (l Line) Pos() Point {
	return l.Start
}

func (l Line) Centroid() Point {
	return toPoint(xy.LinesCentroid(l.toGeom()))
}

func (l *Line) Translate(p Point) {
	l.Start.Translate(p)
	l.End.Translate(p)
}

func (l *Line) Rotate(rad float64) {
	l.RotateAround(rad, l.Centroid())
}

func (l *Line) RotateAround(rad float64, pivot Point) {
	l.Translate(pivot.Invert())
	l.Start.Rotate(rad)
	l.End.Rotate(rad)
	l.Translate(pivot)
}

func (l Line) Draw(g *gg.Context) {
	if l.Start != l.End {
		l.DrawOutline(g)
		l.DrawCenter(g)
	}
}

func (l Line) DrawOutline(g *gg.Context) {
	g.Push()
	defer g.Pop()

	g.SetColor(color.Black)
	g.SetLineWidth(2)
	g.DrawLine(l.Start.X, l.Start.Y, l.End.X, l.End.Y)
	g.Stroke()

	g.DrawCircle(l.Start.X, l.Start.Y, 3)
	g.DrawCircle(l.End.X, l.End.Y, 3)
	g.Fill()
}

func (l Line) DrawCenter(g *gg.Context) {
	g.Push()
	defer g.Pop()
	center := l.Centroid()
	center.Draw(g)
}

func (l Line) toGeom() *geom.LineString {
	ge := geom.NewLineString(geom.XY)
	ge.MustSetCoords([]geom.Coord{{l.Start.X, l.Start.Y}, {l.End.X, l.End.Y}})
	return ge
}
