package geom

import (
	"math"

	"github.com/ctessum/polyclip-go"
)

// MultiPolygon is a holder for multiple related polygons.
type MultiPolygon []Polygon

// Bounds gives the rectangular extents of the MultiPolygon.
func (mp MultiPolygon) Bounds() *Bounds {
	b := NewBounds()
	for _, polygon := range mp {
		b.Extend(polygon.Bounds())
	}
	return b
}

// Area returns the combined area of the polygons in p,
//  assuming that none of the polygons in the p
// overlap and that nested polygons have alternating winding directions.
func (mp MultiPolygon) Area() float64 {
	a := 0.
	for _, pp := range mp {
		a += pp.Area()
	}
	return math.Abs(a)
}

// Intersection returns the area(s) shared by mp and p2.
func (mp MultiPolygon) Intersection(p2 Polygonal) Polygon {
	return mp.op(p2, polyclip.INTERSECTION)
}

// Union returns the combination of mp and p2.
func (mp MultiPolygon) Union(p2 Polygonal) Polygon {
	return mp.op(p2, polyclip.UNION)
}

// XOr returns the area(s) occupied by either mp or p2 but not both.
func (mp MultiPolygon) XOr(p2 Polygonal) Polygon {
	return mp.op(p2, polyclip.XOR)
}

// Difference subtracts p2 from mp.
func (mp MultiPolygon) Difference(p2 Polygonal) Polygon {
	return mp.op(p2, polyclip.DIFFERENCE)
}

func (mp MultiPolygon) op(p2 Polygonal, op polyclip.Op) Polygon {
	var pp polyclip.Polygon
	for _, ppx := range mp.Polygons() {
		pp = append(pp, ppx.toPolyClip()...)
	}
	var pp2 polyclip.Polygon
	for _, pp2x := range p2.Polygons() {
		pp2 = append(pp2, pp2x.toPolyClip()...)
	}
	return polyClipToPolygon(pp.Construct(op, pp2))
}

// Polygons returns the polygons that make up mp.
func (mp MultiPolygon) Polygons() []Polygon {
	return mp
}

// Within calculates whether mp is completely within p. Edges that touch are
// considered to be within. It may not work correctly if mp has holes.
func (mp MultiPolygon) Within(p Polygonal) bool {
	for _, poly := range mp {
		for _, r := range poly {
			for _, pp := range r {
				if !pointInPolygonal(pp, p) {
					return false
				}
			}
		}
	}
	return true
}

// Centroid calculates the centroid of mp, from
// wikipedia: http://en.wikipedia.org/wiki/Centroid#Centroid_of_polygon.
// The polygon can have holes, but each ring must be closed (i.e.,
// p[0] == p[n-1], where the ring has n points) and must not be
// self-intersecting.
// The algorithm will not check to make sure the holes are
// actually inside the outer rings.
func (mp MultiPolygon) Centroid() Point {
	var A, xA, yA float64
	for _, p := range mp {
		for _, r := range p {
			a := area(r)
			cx, cy := 0., 0.
			for i := 0; i < len(r)-1; i++ {
				cx += (r[i].X + r[i+1].X) *
					(r[i].X*r[i+1].Y - r[i+1].X*r[i].Y)
				cy += (r[i].Y + r[i+1].Y) *
					(r[i].X*r[i+1].Y - r[i+1].X*r[i].Y)
			}
			cx /= 6 * a
			cy /= 6 * a
			A += a
			xA += cx * a
			yA += cy * a
		}
	}
	return Point{X: xA / A, Y: yA / A}
}
