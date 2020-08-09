package goeonet

import (
  "testing"
)

func TestGetRecentOpenEventsBasic(t *testing.T) {
  collection, err := GetRecentOpenEvents(1)

  if err != nil {
    t.Error("TestGetRecentOpenEventsBasic: ", err)
  }

  if collection.Title != "EONET Events" && collection.Link != baseEventsUrl {
    t.Error("TestGetRecentOpenEventsBasic: an error has likely occurred while querying the events API")
  }

  if len(collection.Events) > 1 {
    t.Error("TestGetRecentOpenEventsBasic: number of events returned exceeded the limit")
  }
}

func TestGetEventsByDateBasic(t *testing.T) {
  collection, err := GetEventsByDate("2010-01-01", "2020-01-01")

  if err != nil {
    t.Error("TestGetEventsByDateBasic: ", err)
  }

  if len(collection.Events) < 1 {
    t.Error("TestGetEventsByDateBasic: there should be at least some events that occured from 2010-2020")
  }
}

func TestGetEventsBySourceID(t *testing.T) {
  collection, err := GetEventsBySourceID("PDC")

  if err != nil {
    t.Error("TestGetEventsBySourceID: ", err)
  }

  if len(collection.Events) < 1 {
    t.Error("TestGetEventsBySourceID: there should be at least some events whose source is the Pacific Disaster Center (PDC)")
  }
}

func TestGetSourcesBasic(t *testing.T) {
  collection, err := GetSources()

  if err != nil {
    t.Error("TestGetSourcesBasic: ", err)
  }

  if collection.Title != "EONET Event Sources" {
    t.Error("TestGetSourcesBasic: the title returned from the api doesn't match")
  }

  if collection.Link != baseSourcesUrl {
    t.Error("TestGetSourcesBasic: the link returned from the api doesn't match")
  }

  if len(collection.Sources) < 1 {
    t.Error("TestGetSourcesBasic: there should be at least some sources returned by the API")
  }
}

func TestGetCategoriesBasic(t *testing.T) {
  collection, err := GetCategories()

  if err != nil {
    t.Error("TestGetCategoriesBasic: ", err)
  }

  if collection.Title != "EONET Event Categories" {
    t.Error("TestGetCategoriesBasic: the title returned from the api doesn't match")
  }

  if collection.Link != baseCategoriesUrl {
    t.Error("TestGetCategoriesBasic: the link returned from the api doesn't match")
  }
}

func TestGetEventsByCategoryIDLandslides(t *testing.T) {
  collection, err := GetEventsByCategoryID("landslides")

  if err != nil {
    t.Error("TestGetEventsByCategoryIDLandslides: ", err)
  }

  if collection.Title != "EONET Events: Landslides" {
    t.Error("TestGetEventsByCategoryIDLandslides: the title returned from the api doesn't match")
  }

  if collection.Link != baseCategoriesUrl + "/landslides" {
    t.Error("TestGetEventsByCategoryIDLandslides: the link returned from the api doesn't match")
  }
}
