package goeonet

import (
  "encoding/json"
  "errors"
  "net/url"
  "time"
)

const layoutISO = "2006-01-02"

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

type eventsQuery struct {
	source string
	status string
	limit  string
	days   string
	start  string
	end    string
	magID  string
	magMin string
	magMax string
	bbox   string
}

func GetRecentOpenEvents(limit string) (*Collection, error) {
	url := createEventsApiUrl(eventsQuery{limit: limit, status: "open"})

	collection, err := queryEventsApi(url.String())
	if err != nil {
		return nil, err
	}

	return collection, nil
}

func GetRecentClosedEvents(limit string) (*Collection, error) {
	url := createEventsApiUrl(eventsQuery{limit: limit, status: "closed"})

	collection, err := queryEventsApi(url.String())
	if err != nil {
		return nil, err
	}

	return collection, nil
}

func GetEventsByDate(startDate, endDate string) (*Collection, error) {
	if !isValidDate(startDate) {
		return nil, errors.New("the starting date is invalid")
	}

	if endDate != "" && !isValidDate(endDate) {
		return nil, errors.New("the ending date is invalid")
	}

	url := createEventsApiUrl(eventsQuery{start: startDate, end: endDate})

	collection, err := queryEventsApi(url.String())
	if err != nil {
		return nil, err
	}

	return collection, nil
}

func isValidDate(date string) bool {
	_, err := time.Parse(layoutISO, date)
	if err != nil {
		return false
	} else {
		return true
	}
}

func GetEventsBySourceID(sourceID string) (*Collection, error) {
	url := createEventsApiUrl(eventsQuery{source: sourceID})

	collection, err := queryEventsApi(url.String())
	if err != nil {
		return nil, err
	}

	return collection, nil
}

func createEventsApiUrl(query eventsQuery) url.URL {
	u := url.URL {
		Scheme: "https",
		Host: "eonet.sci.gsfc.nasa.gov",
		Path: "/api/v3/events",
	}
	q := u.Query()
	q.Set("source", query.source)
	q.Set("status", query.status)
	q.Set("limit", query.limit)
	q.Set("days", query.days)
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
