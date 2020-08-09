# EONET Client
A client written in Golang for getting information on current natural events provided by https://eonet.sci.gsfc.nasa.gov/
### Installing
- run: `go get github.com/jfitzzg/goeonet` in the console. Must have Golang and Git installed
### Examples
```
// get the 10 most recent open events
collection, _ := goeonet.GetRecentOpenEvents(10)

// get all events that have occurred since January 1st, 2010
collection, _ := goeonet.GetEventsByDate("2010-01-01", "")

// get a list of all the organizations used as sources by EONET
collection, _ := goeonet.GetSources()

// get all events with the specified source ID's
collection, _ := goeonet.GetEventsBySourceID("InciWeb,EO")
```
### License
MIT
