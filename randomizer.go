package randomizer

import (
	"math/rand"

	"github.com/uber/h3-go"
)

type Randomizer struct {
	MinLongtitude float64
	MaxLongtitude float64
	MinLatitude   float64
	MaxLatitude   float64
	H3Resolution  int
}

// NewRandomizer creates a new instance of Randomizer
func NewRandomizer(minLon, maxLon, minLat, maxLat float64) *Randomizer {
	return &Randomizer{
		MaxLatitude:   maxLat,
		MaxLongtitude: maxLon,
		MinLatitude:   minLat,
		MinLongtitude: minLon,
	}
}

// RandomLocation generates a random location
func (r *Randomizer) RandomLocation() (Location, error) {
	return Location{
		Lat: r.getLatitude(),
		Lon: r.getLongtitude(),
	}, nil
}

// RandomH3 generates a random h3 cell ID and returns it with the associated location
func (r *Randomizer) RandomH3(resolution int) (string, Location, error) {
	location, err := r.RandomLocation()

	if err != nil {
		return "", Location{}, err
	}

	geo := h3.GeoCoord{
		Latitude:  location.Lat,
		Longitude: location.Lon,
	}

	index := h3.FromGeo(geo, resolution)

	stringIndex := h3.ToString(index)

	return stringIndex, location, nil

}

func (r *Randomizer) getLongtitude() float64 {

	return r.MinLongtitude + rand.Float64()*(r.MaxLongtitude-r.MinLongtitude)

}

func (r *Randomizer) getLatitude() float64 {

	return r.MinLatitude + rand.Float64()*(r.MaxLatitude-r.MinLatitude)

}
