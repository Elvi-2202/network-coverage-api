package utils

import "math"
import "github.com/yageek/lambertgo"

// HaversineDistance calcule la distance en kilomètres entre deux points géographiques
// spécifiés par leurs latitudes et longitudes en degrés.
func HaversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Rayon de la Terre en kilomètres
	dLat := (lat2 - lat1) * (math.Pi / 180)
	dLon := (lon2 - lon1) * (math.Pi / 180)	
	lat1Rad := lat1 * (math.Pi / 180)
	lat2Rad := lat2 * (math.Pi / 180)	
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1Rad)*math.Cos(lat2Rad)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))	
	return R * c
}

func LambertToGPS(x, y float64) (float64, float64) {
	
	p := lambertgo.Point{X: x, Y: y}

	p.ToWGS84(lambertgo.Lambert93)

	return p.Y, p.X
}
