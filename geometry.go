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

func (*g Geometry) UnmarshalJSON(data []byte) error {
  var object map[string]interface{}
  err := json.Unmarshal(data, &object)
  if err != nil {
    return err
  }

  return decodeGeometry(g, object)
}

func decodeGeometry(g *Geometry, object map[string]interface{}) error {
  mv, ok := object["magnitudeValue"]
  if !ok {
    return errors.New("magnitudeValue property is not defined")
  }

  if s, ok := mv.(string); ok {
    g.MagnitudeValue = s
  } else if f, ok := mv.(float64); ok {
    g.MagnitudeValue = f
  } else {
    return errors.New("magnitudeValue property is neither string nor float64")
  }

  mu, ok := object["magnitudeUnit"]
  if !ok {
    return errors.New("magnitudeUnit property is not defined")
  }

  if s, ok := mv.(string); ok {
    g.MagnitudeUnit = s
  } else {
    return errors.New("magnitudeUnit property is not a string")
  }

  d, ok := object["date"]
  if !ok {
    return errors.New("date property is not defined")
  }

  if s, ok := d.(string); ok {
    g.Date = s
  } else {
    return errors.New("date property is not a string")
  }

  t, ok := object["type"]
  if !ok {
    return errors.New("type property is not defined")
  }

  if s, ok := t.(string); ok {
    g.Type = GeometryType(s)
  } else {
    return errors.New("type property is not a string")
  }

  switch g.Type {
  case GeometryPoint:
    g.Point, err = decodePoint(object["coordinates"])
  case GeometryPolygon:
    g.Polygon, err = decodePolygon(object["coordinates"])
  }

  return nil
}
