package randomizer

type Location struct {
	Lon float64
	Lat float64
}

func (l Location) getLongtitude() float64 {
	return l.Lon
}

func (l Location) getLatitude() float64 {
	return l.Lat
}
