package goeonet

import "testing"

func TestGetRecentOpenEventsBasic(t *testing.T) {
  collection, err := GetRecentOpenEvents("1")

  if err != nil {
    t.Error(err)
  }

  if collection.Title != "EONET Events" && collection.Link != baseEventsUrl {
    t.Error("An error has likely occurred while querying the events API")
  }

  if len(collection.Events) > 1 {
    t.Error("Number of events returned exceeded the limit")
  }
}

func TestGetEventsByDateBasic(t *testing.T) {
  collection, err := GetEventsByDate("2010-01-01", "2020-01-01")

  if err != nil {
    t.Error(err)
  }

  if len(collection.Events) < 1 {
    t.Error("There should be at least some events that occured from 2010-2020")
  }
}

func TestGetEventsByDateBadStartDate(t *testing.T) {
  _, err := GetEventsByDate("01-01-2010", "")

  if err == nil {
    t.Error("An invalid format for the start date was used successfully")
  }
}

func TestGetEventsBySourceID(t *testing.T) {
  collection, err := GetEventsBySourceID("PDC")

  if err != nil {
    t.Error(err)
  }

  if len(collection.Events) < 1 {
    t.Error("There should be at least some events whose source is the Pacific Disaster Center (PDC)")
  }
}
