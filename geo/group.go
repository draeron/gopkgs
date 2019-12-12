package geo

import (
	"github.com/fogleman/gg"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/xy"
)

type Group struct {
	pts  Point
	geos []*Triangle
}

func NewGroup() Group {
	return Group{}
}

func (g *Group) Add(triangle *Triangle) {
	g.geos = append(g.geos, triangle)
}

func (g *Group) SetPos(pos Point) {
	g.pts = pos
}

func (g Group) bounds() *geom.Bounds {
	return g.multiPoly().Bounds()
}

func (g Group) multiPoly() *geom.MultiPolygon {
	mp := geom.NewMultiPolygon(geom.XY)
	for _, tr := range g.geos {
		mp.Push(tr.poly)
	}
	return mp
}

func (g Group) Bounds() Rectangle {
	return ToRect(g.bounds())
}

func (g *Group) Intersect(ge Geometry) bool {
	return intersect(g, ge)
}

func (g Group) Pos() Point {
	return g.pts
}

func (g Group) Centroid() Point {
	return toPoint(xy.MultiPolygonCentroid(g.multiPoly()))
}

func (g *Group) Translate(p Point) {
	for i := 0; i < len(g.geos); i++ {
		g.geos[i].Translate(p)
	}
	//g.pts.Translate(p)
}

func (g *Group) Rotate(rad float64) {
	g.RotateAround(rad, g.Pos())
}

func (g *Group) RotateAround(rad float64, pivot Point) {
	for _, tr := range g.geos {
		tr.RotateAround(rad, pivot)
	}
}

func (g Group) Draw(ctx *gg.Context) {
	g.DrawOutline(ctx)
	g.DrawCenter(ctx)
}

func (g Group) DrawOutline(ctx *gg.Context) {
	for _, tr := range g.geos {
		tr.DrawOutline(ctx)
	}
}

func (g Group) DrawCenter(ctx *gg.Context) {
	for _, tr := range g.geos {
		tr.DrawCenter(ctx)
	}
}
