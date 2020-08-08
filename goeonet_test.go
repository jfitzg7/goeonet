package eonet

import (
  "testing"
)

func TestGetRecentOpenEventsBasic(t *testing.T) {
  eventCollection, err := GetRecentOpenEvents(1)

  if err != nil {
    t.Error("TestGetRecentOpenEventsBasic: ", err)
  }

  if eventCollection.Title != "EONET Events" && eventCollection.Link != baseEventsUrl {
    t.Error("TestGetRecentOpenEventsBasic: an error has likely occurred while querying the events API")
  }

  if len(eventCollection.Events) > 1 {
    t.Error("TestGetRecentOpenEventsBasic: number of events returned exceeded the limit")
  }
}

func TestGetEventsByDateBasic(t *testing.T) {
  eventCollection, err := GetEventsByDate("2010-01-01", "2020-01-01")

  if err != nil {
    t.Error("TestGetEventsByDateBasic: ", err)
  }

  if len(eventCollection.Events) < 1 {
    t.Error("TestGetEventsByDateBasic: there should be at least some events that occured from 2010-2020")
  }
}

func TestGetEventsBySourceID(t *testing.T) {
  eventCollection, err := GetEventsBySourceID("PDC")

  if err != nil {
    t.Error("TestGetEventsBySourceID: ", err)
  }

  if len(eventCollection.Events) < 1 {
    t.Error("TestGetEventsBySourceID: there should be at least some events whose source is the Pacific Disaster Center (PDC)")
  }
}

func TestGetSourcesBasic(t *testing.T) {
  sources, err := GetSources()

  if err != nil {
    t.Error("TestGetSourcesBasic: ", err)
  }

  if len(sources.Sources) < 1 {
    t.Error("TestGetSourcesBasic: there should be at least some sources returned by the API")
  }
}
