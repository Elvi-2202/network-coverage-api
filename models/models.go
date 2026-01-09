package models

type Tower struct {
	Operator string
	Lat, Lon float64
	G2, G3, G4 bool
}


type TechStatus struct {
	G2 bool `json:"2G"`
	G3 bool `json:"3G"`
	G4 bool `json:"4G"`
}