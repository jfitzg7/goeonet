package goeonet

import (
	"errors"
	"net/url"
)

const baseCategoriesUrl = "https://eonet.sci.gsfc.nasa.gov/api/v3/categories"

// Used for specifying the query parameters that can be passed
// to the GetEventsByCategory function. Keep in mind that the
// Category field must be defined if you call GetEventsByCategory.
type CategoriesQueryParameters struct {
	Category string
	Source   string
	Status   string
	Limit    string
	Days     string
}

// Get a list of all of the event categories used by the EONET API
func GetCategories() ([]byte, error) {
	responseData, err := sendRequestToEonetApi(baseCategoriesUrl)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

// Get a list of all the events under a specific category. Remember to always
// assign a value to the Category field in the query parameter.
func GetEventsByCategory(query CategoriesQueryParameters) ([]byte, error) {
	url, err := createCategoriesApiUrl(query)
	if err != nil {
		return nil, err
	}

	responseData, err := sendRequestToEonetApi(url.String())
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func createCategoriesApiUrl(query CategoriesQueryParameters) (*url.URL, error) {
	if query.Category == "" {
		return nil, errors.New("The category must be specified in order to construct the url")
	}

	u := url.URL{
		Scheme: "https",
		Host:   "eonet.sci.gsfc.nasa.gov",
		Path:   "/api/v3/categories/" + query.Category,
	}
	q := u.Query()
	if query.Source != "" {
		q.Set("source", query.Source)
	}
	if query.Status != "" {
		q.Set("status", query.Status)
	}
	if query.Limit != "" {
		q.Set("limit", query.Limit)
	}
	if query.Days != "" {
		q.Set("days", query.Days)
	}
	u.RawQuery = q.Encode()
	return &u, nil
}
