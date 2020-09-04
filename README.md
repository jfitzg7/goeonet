# EONET Client
A client written in Golang for getting information on current natural events provided by https://eonet.sci.gsfc.nasa.gov/
### Installing
To get the latest version run:

`GO111MODULE=on go get github.com/jfitzg7/goeonet`

Must have Golang and Git installed
### Examples
```
// get the 10 most recent open events
response, _ := goeonet.GetEvents(EventsQueryParameters{Limit: 10, Status: "open"})

// get all events that have occurred since January 1st, 2010
response, _ := goeonet.GetEvents(EventsQueryParameters{Start: "2010-01-01"})

// get a list of all the organizations used as sources by EONET
response, _ := goeonet.GetSources()

// get all events in the specified category
response, _ := goeonet.GetEventsByCategory(CategoriesQueryParameters{Category: "wildfires", Source: "InciWeb", Limit: 10, Status: "open"})
```
### Why
This package takes care of all the boilerplate code required to communicate with the EONET API so that you don't have to. All you need to do is pass the query parameters to the functions that you want to use and then you will be able to deal with the JSON response using the parser of your choice.
#### Why no parsing?
I chose not to parse the JSON responses into structs for the user because there are several fields in the API that can have varying types which makes it difficult to parse using the standard encoding/json package. I believe that it would be much easier to use a package that can handle JSON like this with ease, such as [jsonparser](https://github.com/buger/jsonparser) or [gabs](https://github.com/Jeffail/gabs)
### License
MIT
