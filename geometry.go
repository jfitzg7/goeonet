package goeonet

import (
  "encoding/json"
  "errors"
  "fmt"
)

/*
  Inspired by paulmach's go.geojson package. The EONET API
  uses foreign members in their geometry objects, so a different
  implementation must be used for unmarshaling the JSON. Otherwise I
  would just use go.geojson.
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

  return err
}

func decodePoint(data interface{}) ([]float64, error) {
  coords, ok := data.([]interface{})
  if !ok {
    return nil, fmt.Errorf("not a valid point, got %v", data)
  }

  result := make([]float64, 0, len(coords))
  for _, coord := range coords {
    if f, ok := coord.(float64); ok {
      result = append(result, f)
    } else {
      return nil, fmt.Errorf("not a valid coordinate, got %v", coord)
    }
  }

  return result, nil
}

func decodePointSet(data interface{}) ([][]float64, error) {
  points, ok := data.([]interface{})
  if !ok {
    return nil, fmt.Errorf("not a valid set of points, got %v", data)
  }

  result := make([][]float64, 0, len(points))
	for _, point := range points {
		if p, err := decodePoint(point); err == nil {
			result = append(result, p)
		} else {
			return nil, err
		}
	}

	return result, nil
}

func decodePolygon(data interface{}) ([][][]float64, error) {
  sets, ok := data.([]interface{})
	if !ok {
		return nil, fmt.Errorf("not a valid path, got %v", data)
	}

	result := make([][][]float64, 0, len(sets))

	for _, set := range sets {
		if s, err := decodePointSet(set); err == nil {
			result = append(result, s)
		} else {
			return nil, err
		}
	}

	return result, nil
}
