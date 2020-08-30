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

func TestGetLayersBasic(t *testing.T) {
  collection, err := GetLayers()

  if err != nil {
    t.Error(err)
  }

  if collection.Title != "EONET Web Service Layers" {
    t.Error("The title returned from the api doesn't match")
  }

  if collection.Link != baseLayersUrl {
    t.Error("The link returned from the api doesn't match")
  }
}

func TestGetLayersByCategoryIDWildfires(t *testing.T) {
  collection, err := GetLayersByCategoryID("wildfires")

  if err != nil {
    t.Error(err)
  }

  if collection.Categories[0].Title != "Wildfires" {
    t.Error("The title for the wildfires category doesn't match")
  }

  if collection.Categories[0].Id.Id != "8" {
    t.Error("The id for the wildfires category doesn't match")
  }
}
