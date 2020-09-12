[![Go Report Card](https://goreportcard.com/badge/github.com/jfitzg7/goeonet)](https://goreportcard.com/report/github.com/jfitzg7/goeonet)
[![Go Doc Reference](https://godoc.org/github.com/jfitzg7/goeonet?status.svg)](https://godoc.org/github.com/jfitzg7/goeonet)
![GitHub](https://img.shields.io/github/license/jfitzg7/goeonet?color=blue)
# EONET Client
A client written in Golang for getting information on natural events provided by https://eonet.sci.gsfc.nasa.gov/. This package takes care of all the boilerplate code required to communicate with the EONET API so that you don't have to. Just pass the query parameters (when necessary) to the functions you want to use and then handle the JSON response using the parser of your choice.
### Installing
To get the latest version use: `GO111MODULE=on go get github.com/jfitzg7/goeonet@v1.0.0`

### Why no parsing?
I chose not to provide any parsing for the user because there are several fields in the EONET API that can have varying types which makes it difficult to parse the JSON into structs using the standard encoding/json package. I believe it would be better to use a package that can handle dynamic JSON with ease, such as [jsonparser](https://github.com/buger/jsonparser) or [gabs](https://github.com/Jeffail/gabs), so that the user can more easily navigate the responses.
### Query Parameters
The following structs can be passed to the GetEvents() and GetEventsByCategory() functions in order to specify the parameters to be used in the URL query. For more information on the query parameters check out the [EONET API specification](https://eonet.sci.gsfc.nasa.gov/docs/v3)
```go
type EventsQueryParameters struct {
	Source string
	Status string
	Limit  uint
	Days   uint
	Start  string
	End    string
	MagID  string
	MagMin string
	MagMax string
	Bbox   string
}

type CategoriesQueryParameters struct {
	Source   string
	Status   string
	Limit    string
	Days     string
}
```
### Examples
```go
// get the 10 most recent open events
jsonResponse, err := goeonet.GetEvents(goeonet.EventsQueryParameters{Limit: 10, Status: "open"})

// get all events that have occurred since January 1st, 2010
jsonResponse, err := goeonet.GetEvents(goeonet.EventsQueryParameters{Start: "2010-01-01"})

// get all open events in a GeoJSON format
jsonResponse, err := goeonet.GetGeoJsonEvents(goeonet.EventsQueryParameters{Status: "open"})

// get a list of all the organizations used as sources by EONET
jsonResponse, err := goeonet.GetSources()

// get a list of all the categories
jsonResponse, err := goeonet.GetCategories()

// get the 10 most recently closed wildfire events reported by InciWeb
jsonResponse, err := goeonet.GetEventsByCategory("wildfires", goeonet.CategoriesQueryParameters{Source: "InciWeb", Limit: 10, Status: "closed"})

// get a list of all the web service layers
jsonResponse, err := goeonet.GetLayers()

// get a list of the web service layers by category
jsonResponse, err := goeonet.GetLayersByCategory("landslides")

// get a list of all the units used for measuring the magnitudes of specific events
jsonResponse, err := goeonet.GetMagnitudes()
```
