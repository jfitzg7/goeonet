package goeonet

import "testing"

func TestGetSourcesBasic(t *testing.T) {
  collection, err := GetSources()

  if err != nil {
    t.Error(err)
  }

  if collection.Title != "EONET Event Sources" {
    t.Error("The title returned from the api doesn't match")
  }

  if collection.Link != baseSourcesUrl {
    t.Error("The link returned from the api doesn't match")
  }

  if len(collection.Sources) < 1 {
    t.Error("There should be at least some sources returned by the API")
  }
}
