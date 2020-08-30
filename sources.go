package goeonet

import (
  "encoding/json"
)

type Source struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Source string `json:"source"`
	Link   string `json:"link"`
}

type Sources struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Link        string   `json:"link"`
	Sources     []Source `json:"sources"`
}

func GetSources() (*Collection, error) {
	collection, err := querySourcesApi()
	if err != nil {
		return nil, err
	}

	return collection, nil
}

func querySourcesApi() (*Collection, error) {
	responseData, err := sendRequest(baseSourcesUrl)
	if err != nil {
		return nil, err
	}

	var collection Collection

	if err := json.Unmarshal(responseData, &collection); err != nil {
		return nil, err
	}

	return &collection, nil
}
