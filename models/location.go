package models

import (
	"github.com/kellydunn/golang-geo"
)

// Location defines the geospatial location data model.
type Location struct {
	Name      string
	Latitude  float64
	Longitude float64
}

func (l Location) IsNearBy(l2 Location, d float64) bool {
	return l.toGeoPoint().GreatCircleDistance(l2.toGeoPoint()) <= d
}

func (l Location) Distance(l2 Location) float64 {
	return l.toGeoPoint().GreatCircleDistance(l2.toGeoPoint())
}

func (l Location) toGeoPoint() *geo.Point {
	return geo.NewPoint(l.Latitude, l.Longitude)
}
