package geo

import (
	"github.com/draeron/gopkgs/color"
	"github.com/fogleman/gg"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/xy"
)

type Point struct {
	X, Y float64
}

var (
	Zero = Point{0, 0}
)

func NewPointI(x, y int32) Point {
	return Point{float64(x), float64(y)}
}

func (p Point) to3D() Point3D {
	return Point3D{p.X, p.Y, 0}
}

func toPoint(pts geom.Coord) Point {
	return Point{pts.X(), pts.Y()}
}

func (p Point) ToCoord() geom.Coord {
	return geom.Coord{p.X, p.Y}
}

func (p1 *Point) Translate(p Point) {
	p1.X += p.X
	p1.Y += p.Y
}

func (p Point) Bounds() Rectangle {
	return Rectangle{p, 0, 0}
}

func (p Point) bounds() *geom.Bounds {
	b := geom.NewBounds(geom.XY)
	b.SetCoords(p.ToCoord(), p.ToCoord())
	return b
}

func (p Point) Add(p2 Point) Point {
	return p.AddXY(p2.X, p2.Y)
}

func (p Point) AddXY(x, y float64) Point {
	return Point{p.X + x, p.Y + y}
}

func (p *Point) Rotate(rad float64) {
	r := Rotate(*p, rad)
	*p = r
}

func (p *Point) RotateAround(rad float64, pivot Point) {
	*p = RotateAround(*p, pivot, rad)
}

func (p *Point) Intersect(ge Geometry) bool {
	return intersect(p, ge)
}

func (p Point) Pos() Point {
	return p
}

func (p Point) Centroid() Point {
	return p
}

func (p Point) DistanceTo(p2 Point) float64 {
	return xy.Distance(p.ToCoord(), p2.ToCoord())
}

func (p Point) Draw(g *gg.Context) {
	p.DrawCenter(g)
	p.DrawOutline(g)
}

func (p Point) DrawOutline(g *gg.Context) {
	g.Push()
	defer g.Pop()

	g.SetColor(color.Black)
	g.SetLineWidth(1)
	g.DrawPoint(p.X, p.Y, 3)
	g.Stroke()
}

func (p Point) DrawCenter(g *gg.Context) {
	g.Push()
	defer g.Pop()

	g.DrawCircle(p.X, p.Y, 3)
	g.Fill()
}

func (p Point) Invert() Point {
	return Point{-p.X, -p.Y}
}
