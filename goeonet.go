package goeonet

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	layoutISO         = "2006-01-02"
	baseEventsUrl     = "https://eonet.sci.gsfc.nasa.gov/api/v3/events"
	baseCategoriesUrl = "https://eonet.sci.gsfc.nasa.gov/api/v3/categories"
	baseLayersUrl     = "https://eonet.sci.gsfc.nasa.gov/api/v3/layers"
	baseSourcesUrl    = "https://eonet.sci.gsfc.nasa.gov/api/v3/sources"
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

type Source struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Source string `json:"source"`
	Link   string `json:"link"`
}

type Sources struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Link        string   `json:"link"`
	Sources     []Source `json:"sources"`
}

type Collection struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Link        string     `json:"link"`
	Events      []Event    `json:"events,omitempty"`
	Categories  []Category `json:"categories,omitempty"`
	Sources     []Source   `json:"sources,omitempty"`
}

var client = http.Client{Timeout: 5 * time.Second}

func GetSources() (*Collection, error) {
	collection, err := querySourcesApi()
	if err != nil {
		return nil, err
	}

	return collection, nil
}

func querySourcesApi() (*Collection, error) {
	responseData, err := sendRequest(baseSourcesUrl)
	if err != nil {
		return nil, err
	}

	var collection Collection

	if err := json.Unmarshal(responseData, &collection); err != nil {
		return nil, err
	}

	return &collection, nil
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
