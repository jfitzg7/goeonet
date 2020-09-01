package goeonet

import (
  "fmt"
  "net/url"
)

const baseEventsUrl = "https://eonet.sci.gsfc.nasa.gov/api/v3/events"

type EventSource struct {
	Id     string `json:"id"`
	Url    string `json:"url"`
}

type Event struct {
	Id          string        `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Link        string        `json:"link"`
	Closed      string        `json:"closed"`
	Categories  []Category    `json:"categories"`
	Sources     []EventSource `json:"sources"`
	Geometrics  []Geometry    `json:"geometry"`
}

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

func GetEvents(query EventsQueryParameters) (*Collection, error) {
  url := createEventsApiUrl(query)

  collection, err := queryEonetApi(url.String())
  if err != nil {
    return nil, err
  }

  return collection, nil
}

func createEventsApiUrl(query EventsQueryParameters) url.URL {
	u := url.URL {
	   Scheme: "https",
     Host: "eonet.sci.gsfc.nasa.gov",
     Path: "/api/v3/events",
	}
	q := u.Query()
  if query.Source != "" {
	   q.Set("source", query.Source)
  }
  if query.Status != "" {
	   q.Set("status", query.Status)
  }
  if query.Limit > 0 {
	   q.Set("limit", fmt.Sprint(query.Limit))
  }
  if query.Days > 0 {
	   q.Set("days", fmt.Sprint(query.Days))
  }
	if query.Start != "" {
		q.Set("start", query.Start)
    if query.End != "" {
		    q.Set("end", query.End)
    }
	}
  if query.MagID != "" {
    q.Set("magID", query.MagID)
  }
  if query.MagMin != "" {
	  q.Set("magMin", query.MagMin)
  }
  if query.MagMax != "" {
	  q.Set("magMax", query.MagMax)
  }
  if query.Bbox != "" {
	  q.Set("bbox", query.Bbox)
  }
	u.RawQuery = q.Encode()
	return u
}
