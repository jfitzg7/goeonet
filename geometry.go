package goeonet

import (
  "encoding/json"
  "errors"
)

/*
  Inspired by paulmach's go.geojson package. The EONET API
  uses foreign members in their geometry objects, so a custom
  implementation must be used to unmarshal the json properly.
*/

type GeometryType string

const (
  GeometryPoint   GeometryType = "Point"
  GeometryPolygon GeometryType = "Polygon"
)

type Geometry struct {
	MagnitudeValue float64     `json:"magnitudeValue"`
	MagnitudeUnit  string      `json:"magnitudeUnit"`
	Date           string      `json:"date"`
	Type           string      `json:"type"`
	Point    			 []float64
	Polygon        [][][]float64
}
