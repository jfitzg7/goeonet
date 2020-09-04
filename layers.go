package goeonet

import (
  "net/url"
)

const baseLayersUrl = "https://eonet.sci.gsfc.nasa.gov/api/v3/layers"

func GetLayers() ([]byte, error) {
	return GetLayersByCategoryID("")
}

func GetLayersByCategory(categoryID string) ([]byte, error) {
	url := createLayersApiUrl(categoryID)

	responseData, err := sendRequestToEonetApi(url.String())
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func createLayersApiUrl(categoryID string) url.URL {
  var pathExtension string
  if categoryID != "" {
    pathExtension = "/" + categoryID
  } else {
    pathExtension = ""
  }
	u := url.URL {
		Scheme: "https",
		Host: "eonet.sci.gsfc.nasa.gov",
		Path: "/api/v3/layers" + pathExtension,
	}
	return u
}
