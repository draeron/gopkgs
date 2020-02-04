package geo

import (
	"github.com/draeron/gopkgs/color"
	"github.com/fogleman/gg"
	"github.com/twpayne/go-geom"
	"math"
)

type Triangle struct {
	coord       geom.Coord
	poly        *geom.Polygon
	side        float64
	Orientation int32
}

func NewTriangle(pts Point, side int, orientation int32) *Triangle {
	t := Triangle{}

	t.side = float64(side)

	t.coord = pts.ToCoord()
	t.poly = geom.NewPolygon(geom.XY)

	h2 := t.side / 2 * math.Sqrt(3)

	p1 := Point{0, h2 * 2 / 3}
	p2 := Point{t.side / 2, h2 * -1 / 3}
	p3 := Point{-t.side / 2, h2 * -1 / 3}

	radian := Radians(int(orientation))
	p1.Rotate(radian)
	p2.Rotate(radian)
	p3.Rotate(radian)

	p1.Translate(pts)
	p2.Translate(pts)
	p3.Translate(pts)

	t.Orientation = orientation

	ring := geom.NewLinearRing(geom.XY)
	ring.MustSetCoords([]geom.Coord{p1.ToCoord(), p2.ToCoord(), p3.ToCoord(), p1.ToCoord()})
	t.poly.Push(ring)

	return &t
}

func (t Triangle) Bounds() Rectangle {
	return ToRect(t.bounds())
}

func (t Triangle) bounds() *geom.Bounds {
	return t.poly.Bounds()
}

func (t *Triangle) Rotate(rad float64) {
	t.RotateAround(rad, t.Centroid())
}

func (t *Triangle) RotateAround(rad float64, pivot Point) {
	ring := t.poly.LinearRing(0)
	coords := ring.Coords()

	center := t.Centroid()

	for i := 0; i < 4; i++ {
		p := toPoint(coords[i])
		p.RotateAround(rad, pivot)
		coords[i].Set(p.ToCoord())
	}
	t.poly.MustSetCoords([][]geom.Coord{coords})

	center.RotateAround(rad, pivot)
	t.coord = center.ToCoord()
}

func (t *Triangle) Translate(p Point) {
	ctr := toPoint(t.coord)
	ctr.Translate(p)
	t.coord = ctr.ToCoord()

	ring := t.poly.LinearRing(0)
	coords := ring.Coords()

	for i := 0; i < len(coords); i++ {
		coords[i][0] += p.X
		coords[i][1] += p.Y
	}

	t.poly.MustSetCoords([][]geom.Coord{coords})
}

func (t *Triangle) Scale(ratio float64) {
	t.Translate(t.Pos().Invert())

	for i := 0; i < len(t.coord); i++ {
		t.coord[i] -= ratio
	}

	t.Translate(t.Pos())
}

func (t *Triangle) Intersect(ge Geometry) bool {
	return intersect(t, ge)
}

func (t *Triangle) threshold() float64 {
	return t.side / 2 * math.Sqrt(3) * 0.57
}

func (t *Triangle) Pos() Point {
	return toPoint(t.coord)
}

func (t *Triangle) Centroid() Point {
	return toPoint(t.coord)
}

func (t *Triangle) Draw(g *gg.Context) {
	t.DrawOutline(g)
	t.DrawCenter(g)
}

func (t *Triangle) DrawOutline(g *gg.Context) {
	g.Push()
	defer g.Pop()

	for i := 0; i < t.poly.NumCoords(); i++ {
		p1 := t.poly.Coord(i)
		g.LineTo(p1.X(), p1.Y())
	}
	g.FillPreserve()

	g.SetColor(color.Black)
	g.Stroke()
}

func (t *Triangle) DrawCenter(g *gg.Context) {
	g.Push()
	defer g.Pop()

	//ctr := t.Centroid()
	//rgb := color.Black.RGB()
	//rgb.A = 25
	//g.SetColor(rgb)
	//g.DrawCircle(ctr.X, ctr.Y, t.threshold())
	//g.Fill()

	g.SetColor(color.CyanBlue)
	t.Centroid().Draw(g)
}
