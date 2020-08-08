package main

import (
  "fmt"
  "testing"
)

func TestGetRecentOpenEventsBasic(t *testing.T) {
  eventCollection, err := GetRecentOpenEvents(1)

  if eventCollection.Title != "EONET Events" {
    t.Error("TestGetRecentOpenEventsBasic: an error has likely occurred while querying the events API")
  }

  if len(eventCollection.Events) > 1 {
    t.Error("TestGetRecentOpenEventsBasic: number of events returned exceeded the limit")
  }
}
