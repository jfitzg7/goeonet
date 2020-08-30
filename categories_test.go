package goeonet

import "testing"

func TestGetCategoriesBasic(t *testing.T) {
  collection, err := GetCategories()

  if err != nil {
    t.Error(err)
  }

  if collection.Title != "EONET Event Categories" {
    t.Error("The title returned from the api doesn't match")
  }

  if collection.Link != baseCategoriesUrl {
    t.Error("The link returned from the api doesn't match")
  }
}

func TestGetEventsByCategoryIDLandslides(t *testing.T) {
  collection, err := GetEventsByCategoryID("landslides")

  if err != nil {
    t.Error(err)
  }

  if collection.Title != "EONET Events: Landslides" {
    t.Error("The title returned from the api doesn't match")
  }

  if collection.Link != baseCategoriesUrl + "/landslides" {
    t.Error("The link returned from the api doesn't match")
  }
}
