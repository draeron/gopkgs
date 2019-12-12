package geo

import (
	"github.com/fogleman/gg"
	"github.com/twpayne/go-geom"
	"github.com/draeron/gopkg/logger"
)

type Geometry interface {
	Pos() Point
	Centroid() Point
	Intersect(geom Geometry) bool
	Translate(p Point)
	Rotate(rad float64)
	RotateAround(rad float64, anchor Point)
	Bounds() Rectangle

	Draw(g *gg.Context)
	DrawOutline(g *gg.Context)
	DrawCenter(g *gg.Context)

	bounds() *geom.Bounds
}

var log = logger.New("geometry")
