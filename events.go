package goeonet

import (
  "encoding/json"
  "fmt"
  "net/url"
)

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

type EventsQuery struct {
	source string
	status string
	limit  uint
	days   uint
	start  string
	end    string
	magID  string
	magMin string
	magMax string
	bbox   string
}

func GetEvents(query EventsQuery) (*Collection, error) {
  url := createEventsApiUrl(query)

  collection, err := queryEventsApi(url.String())
  if err != nil {
    return nil, err
  }

  return collection, nil
}

func createEventsApiUrl(query EventsQuery) url.URL {
	u := url.URL {
		Scheme: "https",
		Host: "eonet.sci.gsfc.nasa.gov",
		Path: "/api/v3/events",
	}
	q := u.Query()
	q.Set("source", query.source)
	q.Set("status", query.status)
  if query.limit > 0 {
	   q.Set("limit", fmt.Sprint(query.limit))
  }
  if query.days > 0 {
	   q.Set("days", fmt.Sprint(query.days))
  }
	if query.start != "" {
		q.Set("start", query.start)
		q.Set("end", query.end)
	}
	q.Set("magID", query.magID)
	q.Set("magMin", query.magMin)
	q.Set("magMax", query.magMax)
	q.Set("bbox", query.bbox)
	u.RawQuery = q.Encode()
	return u
}

func queryEventsApi(url string) (*Collection, error) {
	responseData, err := sendRequest(url)
	if err != nil {
		return nil, err
	}

	var collection Collection

	if err := json.Unmarshal(responseData, &collection); err != nil {
		return nil, err
	}

	return &collection, nil
}
