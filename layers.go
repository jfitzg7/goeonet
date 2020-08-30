package goeonet

import (
  "encoding/json"
  "net/url"
)

type Parameter struct {
	TILEMATRIXSET string `json:"TILEMATRIXSET"`
	FORMAT        string `json:"FORMAT"`
}

type Layer struct {
	Name          string      `json:"name"`
	ServiceUrl    string      `json:"serviceUrl"`
	ServiceTypeId string      `json:"serviceTypeId"`
	Parameters    []Parameter `json:"parameters"`
}

type Layers struct {
	Link   string
	Layers []Layer
}

func (l *Layers) UnmarshalJSON(data []byte) error {
	if string(data)[0] == 91 { // check if the first character is '['
		var layers []Layer
		err := json.Unmarshal(data, &layers)
		if err != nil {
			return err
		} else {
			l.Layers = layers
			return nil
		}
	} else {
		l.Link = string(data)
		return nil
	}
}

func GetLayers() (*Collection, error) {
	collection, err := queryLayersApi(baseLayersUrl)
	if err != nil {
		return nil, err
	}

	return collection, nil
}

func GetLayersByCategoryID(categoryID string) (*Collection, error) {
	url := createLayersApiUrl(categoryID)

	collection, err := queryLayersApi(url.String())
	if err != nil {
		return nil, err
	}

	return collection, nil
}

func createLayersApiUrl(categoryID string) url.URL {
	u := url.URL {
		Scheme: "https",
		Host: "eonet.sci.gsfc.nasa.gov",
		Path: "/api/v3/layers/" + categoryID,
	}
	return u
}

func queryLayersApi(url string) (*Collection, error) {
	responseData, err := sendRequest(url)
	if err != nil {
		return nil, err
	}

	var collection Collection

	if err := json.Unmarshal(responseData, &collection); err != nil {
		return nil, err
	}

	return &collection, nil
}
