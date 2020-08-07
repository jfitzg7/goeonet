package main

import (
  "encoding/json"
  "fmt"
  "log"
  "io/ioutil"
  "net/http"
  "time"
)

const (
  baseUrl = "https://eonet.sci.gsfc.nasa.gov/api/v3/events"
)

type Category struct {
  Id    string `json:"id"`
  Title string `json:"title"`
}

type Source struct {
  Id  string `json:"id"`
  Url string `json:"url"`
}

type Geometry struct {
  MagnitudeValue string    `json:"magnitudeValue"`
  MagnitudeUnit  string    `json:"magnitudeUnit"`
  Date           string    `json:"date"`
  Type           string    `json:"type"`
  Coordinates    []float64 `json:"coordinates"`
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
  eventCollection, err := GetRecentOpenEvents(10)

  if err != nil {
    log.Fatal("GetRecentOpenEvents: ", err)
  }

  for _, event := range eventCollection.Events {
    fmt.Printf("ID: %s\nTitle: %s\nSources:\n", event.Id, event.Title)
    for _, source := range event.Sources {
      fmt.Printf("\tURL: %s\n", source.Url)
    }
    fmt.Printf("Coordinates: %f, %f\n\n", event.Geometrics[0].Coordinates[0], event.Geometrics[0].Coordinates[1])
  }
}

func GetRecentOpenEvents(limit int) (*EventCollection, error){
  limitParam := fmt.Sprintf("limit=%d", limit)
  request, _ := http.NewRequest("GET", baseUrl + "?status=open&" + limitParam, nil)

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
