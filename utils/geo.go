package utils

import (
	"math"
	"github.com/yageek/lambertgo"
)

func Distance(x, y float64) float64 {
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}
func ToRadian(degree float64) float64 {
	return degree * (math.Pi / 180)
}

func LambertToGPS(x, y float64) (float64, float64) {
	p := lambertgo.Point{X: x, Y: y}
	
	p.ToWGS84(lambertgo.Lambert93)

	return p.Y, p.X
}	
