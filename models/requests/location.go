package requests

import "github.com/timoth-y/chainmetric-core/models"

// LocationQuery defines query for geo location.
type LocationQuery struct {
	// GeoPoint is used to specify query geo location point.
	GeoPoint models.Location `json:"location"`
	// Distance is used to specify satisfactory distance between target and query locations.
	Distance float64         `json:"distance"`
}
