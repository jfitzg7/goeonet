package main

import (
  "encoding/json"
  "fmt"
  "log"
  "io/ioutil"
  "net/http"
  "time"
)

type Source struct {
  Id  string `json:"id"`
  Url string `json:"url"`
}

type Event struct {
  Id      string   `json:"id"`
  Title   string   `json:"title"`
  Sources []Source `json:"sources"`
}

type EventCollection struct {
  Events []Event `json:"events"`
}

var client = http.Client{Timeout: 5 * time.Second}

func main() {
  eventCollection, err := GetRecentOpenEvents()

  if err != nil {
    log.Fatal("GetRecentOpenEvents: ", err)
  }

  for _, event := range eventCollection.Events {
    fmt.Printf("ID: %s\nTitle: %s\nSources:\n", event.Id, event.Title)
    for _, source := range event.Sources {
      fmt.Printf("\tURL: %s\n", source.Url)
    }
    fmt.Println()
  }
}

func GetRecentOpenEvents() (*EventCollection, error){
   request, err := http.NewRequest("GET", "https://eonet.sci.gsfc.nasa.gov/api/v3/events?status=open&limit=10", nil)

   if err != nil {
     return nil, err
   }

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
