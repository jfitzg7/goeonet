package goeonet

import "testing"

func TestGetRecentOpenEvents(t *testing.T) {
  collection, err := GetEvents(EventsQuery{status: "open", limit: 1})

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
  collection, err := GetEvents(EventsQuery{status: "closed", limit: 1})

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
  collection, err := GetEvents(EventsQuery{start: "2010-01-01", end: "2020-01-01"})

  if err != nil {
    t.Error(err)
  }

  if len(collection.Events) < 1 {
    t.Error("There should be at least some events that occured from 2010-2020")
  }
}

func TestGetEventsBySourceID(t *testing.T) {
  collection, err := GetEvents(EventsQuery{source: "PDC"})

  if err != nil {
    t.Error(err)
  }

  if len(collection.Events) < 1 {
    t.Error("There should be at least some events whose source is the Pacific Disaster Center (PDC)")
  }
}
