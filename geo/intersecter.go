package geo

import (
	"fmt"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/xy"
	"math"
	"reflect"
)

type intersectFunc func(Geometry, Geometry) bool

var dispatch = make(map[string]intersectFunc)

func init() {

	pts := reflect.TypeOf(Point{})
	line := reflect.TypeOf(Line{})
	rect := reflect.TypeOf(Rectangle{})
	triangle := reflect.TypeOf(Triangle{})
	circle := reflect.TypeOf(Circle{})
	group := reflect.TypeOf(Group{})
	arc := reflect.TypeOf(Arc{})

	addIntersect(pts, pts, intersectPointPoint)
	addIntersect(line, pts, intersectLinePoint)
	addIntersect(line, line, intersectLineLine)

	addIntersect(triangle, pts, intersectTrianglePoint)
	addIntersect(triangle, line, intersectTriangleLine)
	addIntersect(triangle, triangle, intersectBounds) // TODO intersectTriangleTriangle
	addIntersect(triangle, circle, intersectTriangleCircle)
	addIntersect(triangle, arc, intersectTriangleArc)
	addIntersect(triangle, rect, intersectBounds)  // TODO intersectTriangleRect
	addIntersect(triangle, group, intersectBounds) // TODO intersectTriangleGroup

	addIntersect(circle, pts, intersectBounds)    // todo : intersectCirclePoint
	addIntersect(circle, line, intersectBounds)   // todo : intersectCircleLine
	addIntersect(circle, circle, intersectBounds) // todo : intersectCircleCircle

	addIntersect(arc, arc, intersectBounds)    // todo : intersectArcPoints
	addIntersect(arc, pts, intersectBounds)    // todo : intersectArcPoints
	addIntersect(arc, line, intersectBounds)   // todo : intersectArcLine
	addIntersect(arc, circle, intersectBounds) // todo : intersectArcCircle

	addIntersect(rect, pts, intersectBounds)    // todo : intersectRectPoint
	addIntersect(rect, line, intersectBounds)   // todo : intersectRectLine
	addIntersect(rect, circle, intersectBounds) // todo: intersectRectCircle
	addIntersect(rect, rect, intersectBounds)   // todo: intersectRectRect

	addIntersect(group, pts, intersectGroupGeometry)
	addIntersect(group, line, intersectGroupGeometry)
	addIntersect(group, circle, intersectGroupGeometry)
	addIntersect(group, rect, intersectGroupGeometry)
	addIntersect(group, group, intersectGroupGeometry)
}

func intersect(g1 Geometry, g2 Geometry) bool {

	key := makeKey(typeOf(g1), typeOf(g2))

	fct, ok := dispatch[key]

	if !ok {
		//log.Sugar().Error("unsupported intersection:", key)
		return false
	}

	return fct(g1, g2)
}

func typeOf(g Geometry) reflect.Type {
	switch c := g.(type) {
	case *Point:
		return reflect.TypeOf(*c)
	case *Line:
		return reflect.TypeOf(*c)
	case *Triangle:
		return reflect.TypeOf(*c)
	case *Group:
		return reflect.TypeOf(*c)
	case *Circle:
		return reflect.TypeOf(*c)
	case *Rectangle:
		return reflect.TypeOf(*c)
	case *Arc:
		return reflect.TypeOf(*c)
	default:
		return nil
	}
}

func makeKey(t1 reflect.Type, t2 reflect.Type) string {
	return fmt.Sprintf("%s+%s", t1.Name(), t2.Name())
}

func addIntersect(t1 reflect.Type, t2 reflect.Type, fct intersectFunc) {
	key := makeKey(t1, t2)
	dispatch[key] = fct
	if t1 != t2 {
		key = makeKey(t2, t1)
		dispatch[key] = flipIntersect(fct)
	}
}

func flipIntersect(fct intersectFunc) intersectFunc {
	return func(g1 Geometry, g2 Geometry) bool {
		return fct(g2, g1)
	}
}

func intersectBounds(g1 Geometry, g2 Geometry) bool {
	return g1.bounds().Overlaps(geom.XY, g2.bounds())
}

func intersectPointPoint(g1 Geometry, g2 Geometry) bool {
	p1 := g1.(*Point)
	p2 := g2.(*Point)
	return xy.Distance(p1.ToCoord(), p2.ToCoord()) < math.SmallestNonzeroFloat64
}

func intersectLinePoint(g1 Geometry, g2 Geometry) bool {
	l1 := g1.(*Line)
	p1 := g2.(*Point)

	return xy.DistanceFromPointToLine(p1.ToCoord(), l1.Start.ToCoord(), l1.End.ToCoord()) < math.SmallestNonzeroFloat64
}

func intersectLineLine(g1 Geometry, g2 Geometry) bool {
	l1 := g1.(*Line)
	l2 := g2.(*Line)
	return l1.bounds().Overlaps(geom.XY, l2.bounds())
}

func intersectTrianglePoint(g1 Geometry, g2 Geometry) bool {
	t1 := g1.(*Triangle)
	p1 := g2.(*Point)

	//p0 := toPoint(tr.poly.Coord(0))
	//p1 := toPoint(tr.poly.Coord(1))
	//p2 := toPoint(tr.poly.Coord(2))
	//
	//s := (p0.Y*p2.X - p0.X*p2.Y + (p2.Y-p0.Y)*p.X + (p0.X-p2.X)*p.Y)
	//t := (p0.X*p1.Y - p0.Y*p1.X + (p0.Y-p1.Y)*p.X + (p1.X-p0.X)*p.Y)
	//
	//if s <= math.SmallestNonzeroFloat64 || t <= math.SmallestNonzeroFloat64 {
	//	return false
	//}
	//
	//var A = (-p1.Y*p2.X + p0.Y*(-p1.X+p2.X) + p0.X*(p1.Y-p2.Y) + p1.X*p2.Y)
	//
	//return (s + t) < A

	return xy.Distance(p1.ToCoord(), t1.coord) < t1.threshold()
}

func intersectTriangleLine(g1 Geometry, g2 Geometry) bool {
	t := g1.(*Triangle)
	l := g2.(*Line)

	ring := t.poly.LinearRing(0)

	p0 := ring.Coord(0)
	p1 := ring.Coord(1)
	p2 := ring.Coord(2)

	s := l.Start.ToCoord()
	e := l.End.ToCoord()

	if xy.DistanceFromLineToLine(p0, p1, s, e) < math.SmallestNonzeroFloat64 ||
		xy.DistanceFromLineToLine(p1, p2, s, e) < math.SmallestNonzeroFloat64 ||
		xy.DistanceFromLineToLine(p2, p0, s, e) < math.SmallestNonzeroFloat64 {

		dist := xy.DistanceFromPointToLine(t.coord, l.Start.ToCoord(), l.End.ToCoord())
		return dist < t.threshold()
	}

	return false
}

func intersectTriangleCircle(g1 Geometry, g2 Geometry) bool {
	t := g1.(*Triangle)
	c := g2.(*Circle)

	dist := xy.Distance(t.coord, c.Center.ToCoord())

	return math.Abs(dist-c.Radius) < t.threshold()
}

func intersectTriangleArc(g1, g2 Geometry) bool {
	t := g1.(*Triangle)
	a := g2.(*Arc)

	if a.Width < TwoPi/360*2 {
		return false
	}

	if !a.Circle.Intersect(t) {
		return false
	}

	// test for one end
	pts := a.Center.AddXY(a.Radius, 0)
	pts.RotateAround(a.Angle, a.Center)
	if t.Intersect(&pts) {
		return true
	}
	pts.RotateAround(a.Width, a.Center)
	if t.Intersect(&pts) {
		return true
	}

	angle := xy.AngleBetweenOriented(a.Center.AddXY(a.Radius, 0).ToCoord(), a.Center.ToCoord(), t.Centroid().ToCoord())

	angle = Normalize(angle)

	if a.Angle+a.Width <= TwoPi {
		return angle >= a.Angle && angle <= a.Angle+a.Width
	} else {
		if angle >= a.Angle && angle < TwoPi {
			return true
		} else if angle > 0 && angle <= Normalize(a.Angle+a.Width) {
			return true
		}
	}

	return false
}

func intersectGroupGeometry(g1, g2 Geometry) bool {
	group := g1.(*Group)

	for _, tr := range group.geos {
		if tr.Intersect(g2) {
			return true
		}
	}
	return false

}
