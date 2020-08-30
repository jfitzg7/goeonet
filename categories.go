package goeonet

import (
  "encoding/json"
  "net/url"
)

type CategoryID struct {
	Id string
}

// Convert everything to string since the category id can be either a number or string
func (c *CategoryID) UnmarshalJSON(data []byte) error {
	c.Id = string(data)
	return nil
}

type Category struct {
	Id          CategoryID `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Link        string `json:"link,omitempty"`
	Description string `json:"description,omitempty"`
	Layers      Layers `json:"layers,omitempty"`
}

type categoriesQuery struct {
	category string
	source   string
	status   string
	limit    string
	days     string
}

func GetCategories() (*Collection, error) {
	collection, err := queryCategoriesApi(baseCategoriesUrl)
	if err != nil {
		return nil, err
	}

	return collection, nil
}

func GetEventsByCategoryID(categoryID string) (*Collection, error) {
	url := createCategoriesApiUrl(categoriesQuery{category: categoryID})

	collection, err := queryCategoriesApi(url.String())
	if err != nil {
		return nil, err
	}

	return collection, nil
}

func createCategoriesApiUrl(query categoriesQuery) url.URL {
	u := url.URL {
		Scheme: "https",
		Host: "eonet.sci.gsfc.nasa.gov",
		Path: "/api/v3/categories/" + query.category,
	}
	q := u.Query()
	q.Set("source", query.source)
	q.Set("status", query.status)
	q.Set("limit", query.limit)
	q.Set("days", query.days)
	u.RawQuery = q.Encode()
	return u
}

func queryCategoriesApi(url string) (*Collection, error) {
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
