package randomizer

import (
	"testing"
)

const (
	MinLongtitude = 81.15600585937499
	MaxLongtitude = 80.057373046875
	MinLatitude   = 6.560931820902541
	MaxLatitude   = 8.189742344383703
)

func TestNewRandomizer(t *testing.T) {
	var _ *Randomizer = NewRandomizer(MinLongtitude, MaxLongtitude, MinLatitude, MaxLatitude)
}

func TestRandomLocation(t *testing.T) {
	r := NewRandomizer(MinLongtitude, MaxLongtitude, MinLatitude, MaxLatitude)
	location, _ := r.RandomLocation()

	if location.getLatitude() == 0 || location.getLongtitude() == 0 {
		t.Error("TestRandomLocation: Invalid locations received")
	}

}

func TestRandomH3(t *testing.T) {
	r := NewRandomizer(MinLongtitude, MaxLongtitude, MinLatitude, MaxLatitude)
	h3Cell, location, err := r.RandomH3(8)
	if err != nil {
		t.Error(err)
	}

	if h3Cell == "" {
		t.Error("Invalid H3 Cell returned")
	}

	if location.getLatitude() == 0 || location.getLongtitude() == 0 {
		t.Error("TestRandomLocation: Invalid locations received")
	}

}
