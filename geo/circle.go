package geo

import (
	"github.com/draeron/gopkg/color"
	"github.com/fogleman/gg"
	"github.com/twpayne/go-geom"
)

type Circle struct {
	Center Point
	Radius float64
}

func (c Circle) Pos() Point {
	return c.Center
}

func (c Circle) Centroid() Point {
	return c.Center
}

func (c *Circle) Intersect(geom Geometry) bool {
	return intersect(c, geom)
}

func (c *Circle) Translate(p Point) {
	c.Center.Translate(p)
}

func (Circle) Rotate(rad float64) {
	// nothing to do!
}

func (c *Circle) RotateAround(rad float64, anchor Point) {
	c.Center.RotateAround(rad, anchor)
}

func (c Circle) Bounds() Rectangle {
	return Rectangle{
		c.Center.AddXY(-c.Radius, -c.Radius),
		c.Radius * 2,
		c.Radius * 2,
	}
}

func (c Circle) Draw(g *gg.Context) {
	c.DrawOutline(g)
}

func (c Circle) DrawOutline(g *gg.Context) {
	g.Push()

	g.SetColor(color.Black)
	g.SetLineWidth(4)
	g.DrawCircle(c.Center.X, c.Center.Y, c.Radius)
	g.Stroke()

	g.Pop()
}

func (c Circle) DrawCenter(g *gg.Context) {
	c.Center.Draw(g)
}

func (c Circle) bounds() *geom.Bounds {
	bounds := geom.NewBounds(geom.XY)
	tl := c.Center.AddXY(-c.Radius, -c.Radius)
	lr := c.Center.AddXY(c.Radius, c.Radius)
	bounds.SetCoords(tl.ToCoord(), lr.ToCoord())
	return bounds
}
