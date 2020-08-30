package goeonet

import "encoding/json"

type Geometry struct {
	MagnitudeValue float64     `json:"magnitudeValue"`
	MagnitudeUnit  string      `json:"magnitudeUnit"`
	Date           string      `json:"date"`
	Type           string      `json:"type"`
	Point    			 []float64
	Polygon        [][][]float64
}
