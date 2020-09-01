package goeonet

import (
  "bytes"
  "io/ioutil"
  "testing"
  "net/http"
  "time"

  "github.com/golang/mock/gomock"
  "github.com/jfitzg7/goeonet/mocks"
)

func TestGetEvents(t *testing.T) {
  mockCtrl := gomock.NewController(t)

  mockHTTPClient := mocks.NewMockHTTPClient(mockCtrl)
  client = mockHTTPClient

  request, _ := http.NewRequest("GET", "https://eonet.sci.gsfc.nasa.gov/api/v3/events?limit=1", nil)

  response := &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte("")))}

  mockHTTPClient.EXPECT().Do(gomock.Eq(request)).Return(response, nil).Times(1)

  GetEvents(EventsQueryParameters{Limit: 1})

  client = &http.Client{Timeout: 5 * time.Second}
}

func TestGetRecentOpenEvents(t *testing.T) {
  collection, err := GetEvents(EventsQueryParameters{Status: "open", Limit: 1})

  if err != nil {
    t.Error(err)
  }

  if collection.Title != "EONET Events" && collection.Link != baseEventsUrl {
    t.Error("An error has likely occurred while querying the events API")
  }

  // There might be 0 recent open events, so checking for != 1 won't work
  if len(collection.Events) > 1 {
    t.Error("Number of open events returned exceeded the limit")
  }
}

func TestGetRecentClosedEvents(t *testing.T) {
  collection, err := GetEvents(EventsQueryParameters{Status: "closed", Limit: 1})

  if err != nil {
    t.Error(err)
  }

  if collection.Title != "EONET Events" && collection.Link != baseEventsUrl {
    t.Error("An error has likely occurred while querying the events API")
  }

  // There should always be at least 1 closed event returned
  if len(collection.Events) != 1 {
    t.Error("Number of closed events returned does not match the specified limit")
  }
}

func TestGetEventsByDateBasic(t *testing.T) {
  collection, err := GetEvents(EventsQueryParameters{Start: "2010-01-01", End: "2020-01-01"})

  if err != nil {
    t.Error(err)
  }

  if len(collection.Events) < 1 {
    t.Error("There should be at least some events that occured from 2010-2020")
  }
}

func TestGetEventsBySourceID(t *testing.T) {
  collection, err := GetEvents(EventsQueryParameters{Source: "PDC"})

  if err != nil {
    t.Error(err)
  }

  if len(collection.Events) < 1 {
    t.Error("There should be at least some events whose source is the Pacific Disaster Center (PDC)")
  }
}
