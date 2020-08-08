package goeonet

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	layoutISO         = "2006-01-02"
	baseEventsUrl     = "https://eonet.sci.gsfc.nasa.gov/api/v3/events"
	baseCategoriesUrl = "https://eonet.sci.gsfc.nasa.gov/api/v3/categories"
	baseLayersUrl     = "https://eonet.sci.gsfc.nasa.gov/api/v3/layers"
	baseSourcesUrl    = "https://eonet.sci.gsfc.nasa.gov/api/v3/sources"
)

type Category struct {
	Id    string `json:"id"`
	Title string `json:"title"`
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

type Coordinates struct {
	Coordinates [][]float64
}

func (c *Coordinates) UnmarshalJSON(data []byte) error {
	dataString := strings.Replace(string(data), " ", "", -1)
	dataString = strings.Replace(dataString, "],", "", -1)
	dataString = strings.Replace(dataString, "]", "", -1)
	dataString = strings.Replace(dataString, "[[", "", -1)
	coordinates := make([][]float64, 0)
	for _, coords := range strings.Split(dataString[1:], "[") {
		coordArr := make([]float64, 0)
		split := strings.Split(coords, ",")
		coord1, _ := strconv.ParseFloat(split[0], 64)
		coord2, _ := strconv.ParseFloat(split[1], 64)
		coordArr = append(coordArr, coord1)
		coordArr = append(coordArr, coord2)
		coordinates = append(coordinates, coordArr)
	}
	c.Coordinates = coordinates
	return nil
}

type EventSource struct {
	Id     string `json:"id"`
	Url    string `json:"url"`
}

type Geometry struct {
	MagnitudeValue float64     `json:"magnitudeValue"`
	MagnitudeUnit  string      `json:"magnitudeUnit"`
	Date           string      `json:"date"`
	Type           string      `json:"type"`
	Coordinates    Coordinates `json:"coordinates"`
}

type Event struct {
	Id          string        `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Link        string        `json:"link"`
	Closed      string        `json:"closed"`
	Categories  []Category    `json:"categories"`
	Sources     []EventSource `json:"sources"`
	Geometrics  []Geometry    `json:"geometry"`
}

type EventCollection struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Link        string  `json:"link"`
	Events      []Event `json:"events"`
}

var client = http.Client{Timeout: 5 * time.Second}

func GetRecentOpenEvents(limit uint32) (*EventCollection, error) {
	query := fmt.Sprintf("?status=open&limit=%d", limit)

	eventCollection, err := queryEventsApi(query)
	if err != nil {
		return nil, err
	}

	return eventCollection, nil
}

func GetEventsByDate(startDate, endDate string) (*EventCollection, error) {
	if !isValidDate(startDate) {
		return nil, errors.New("the starting date is invalid")
	}

	if endDate != "" && !isValidDate(endDate) {
		return nil, errors.New("the ending date is invalid")
	}

	query := fmt.Sprintf("?start=%s", startDate)

	if endDate != "" {
		query = query + "&end=" + endDate
	}

	eventCollection, err := queryEventsApi(query)
	if err != nil {
		return nil, err
	}

	return eventCollection, nil
}

func isValidDate(date string) bool {
	_, err := time.Parse(layoutISO, date)
	if err != nil {
		return false
	} else {
		return true
	}
}

func GetEventsBySourceID(sourceID string) (*EventCollection, error) {
	query := fmt.Sprintf("?source=%s", sourceID)

	eventCollection, err := queryEventsApi(query)
	if err != nil {
		return nil, err
	}

	return eventCollection, nil
}

func queryEventsApi(query string) (*EventCollection, error) {
	request, _ := http.NewRequest("GET", baseEventsUrl + query, nil)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var eventCollection EventCollection

	if err := json.Unmarshal(responseData, &eventCollection); err != nil {
		return nil, err
	}

	return &eventCollection, nil
}

func GetSources() (*Sources, error) {
	sources, err := querySourcesApi()
	if err != nil {
		return nil, err
	}

	return sources, nil
}

func querySourcesApi() (*Sources, error) {
	request, _ := http.NewRequest("GET", baseSourcesUrl, nil)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var sources Sources

	if err := json.Unmarshal(responseData, &sources); err != nil {
		return nil, err
	}

	return &sources, nil
}
