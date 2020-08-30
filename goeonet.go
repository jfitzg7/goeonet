package goeonet

import (
	"io/ioutil"
	"net/http"
	"time"
)

const (
	layoutISO         = "2006-01-02"
	baseEventsUrl     = "https://eonet.sci.gsfc.nasa.gov/api/v3/events"
	baseCategoriesUrl = "https://eonet.sci.gsfc.nasa.gov/api/v3/categories"
	baseLayersUrl     = "https://eonet.sci.gsfc.nasa.gov/api/v3/layers"
	baseSourcesUrl    = "https://eonet.sci.gsfc.nasa.gov/api/v3/sources"
)

type Collection struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Link        string     `json:"link"`
	Events      []Event    `json:"events,omitempty"`
	Categories  []Category `json:"categories,omitempty"`
	Sources     []Source   `json:"sources,omitempty"`
}

var client = http.Client{Timeout: 5 * time.Second}

func sendRequest(url string) ([]byte, error) {
	request, _ := http.NewRequest("GET", url, nil)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
