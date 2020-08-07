package main

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
)

type Category struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type Source struct {
	Id  string `json:"id"`
	Url string `json:"url"`
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

type Geometry struct {
	MagnitudeValue float64     `json:"magnitudeValue"`
	MagnitudeUnit  string      `json:"magnitudeUnit"`
	Date           string      `json:"date"`
	Type           string      `json:"type"`
	Coordinates    Coordinates `json:"coordinates"`
}

type Event struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Link        string     `json:"link"`
	Closed      string     `json:"closed"`
	Categories  []Category `json:"categories"`
	Sources     []Source   `json:"sources"`
	Geometrics  []Geometry `json:"geometry"`
}

type EventCollection struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Link        string  `json:"link"`
	Events      []Event `json:"events"`
}

var client = http.Client{Timeout: 5 * time.Second}

func main() {
	eventCollection, err := GetEventsByDate("2006-01-02", "")

	if err != nil {
		log.Fatal("GetRecentOpenEvents: ", err)
	}

	for _, event := range eventCollection.Events {
		fmt.Printf("ID: %s\nTitle: %s\nSources:\n", event.Id, event.Title)
		for _, source := range event.Sources {
			fmt.Printf("\tURL: %s\n", source.Url)
		}
		for _, geometry := range event.Geometrics {
			for _, coords := range geometry.Coordinates.Coordinates {
				fmt.Printf("\tCoordinates: %f, %f\n", coords[0], coords[1])
			}
		}
	}
}

func GetRecentOpenEvents(limit int) (*EventCollection, error) {
	url := fmt.Sprintf("%s?status=open&limit=%d", baseEventsUrl, limit)

	responseData, err := queryApi(url)

	if err != nil {
		return nil, err
	}

	var eventCollection EventCollection

	if err := json.Unmarshal(responseData, &eventCollection); err != nil {
		return nil, err
	}

	return &eventCollection, nil
}

func GetEventsByDate(startDate, endDate string) (*EventCollection, error) {
	if !isValidDate(startDate) {
		return nil, errors.New("the starting date is invalid")
	}

	if endDate != "" && !isValidDate(endDate) {
		return nil, errors.New("the ending date is invalid")
	}

	url := fmt.Sprintf("%s?start=%s", baseEventsUrl, startDate)

	if endDate != "" {
		url = url + "&end=" + endDate
	}

	responseData, err := queryApi(url)

	if err != nil {
		return nil, err
	}

	var eventCollection EventCollection

	if err := json.Unmarshal(responseData, &eventCollection); err != nil {
		return nil, err
	}

	return &eventCollection, nil
}

func isValidDate(date string) bool {
	_, err := time.Parse(layoutISO, date)
	if err != nil {
		return false
	} else {
		return true
	}
}

func queryApi(url string) ([]byte, error) {
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
