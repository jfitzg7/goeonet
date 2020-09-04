package goeonet

import (
	"net/url"
)

const baseLayersUrl = "https://eonet.sci.gsfc.nasa.gov/api/v3/layers"

// GetLayers gets a list of all the web service layers
func GetLayers() ([]byte, error) {
	return GetLayersByCategory("")
}

// GetLayersByCategory gets a list of the web service layers by category
func GetLayersByCategory(category string) ([]byte, error) {
	url := createLayersApiUrl(category)

	responseData, err := sendRequestToEonetApi(url.String())
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func createLayersApiUrl(category string) url.URL {
	var pathExtension string
	if category != "" {
		pathExtension = "/" + category
	} else {
		pathExtension = ""
	}
	u := url.URL{
		Scheme: "https",
		Host:   "eonet.sci.gsfc.nasa.gov",
		Path:   "/api/v3/layers" + pathExtension,
	}
	return u
}
