package goeonet

import (
	"net/url"
)

const baseCategoriesUrl = "https://eonet.sci.gsfc.nasa.gov/api/v3/categories"

// Used for specifying the query parameters that can be passed to the
// GetEventsByCategory function. More information on the query parameters
// can be found at https://eonet.sci.gsfc.nasa.gov/docs/v3
type CategoriesQueryParameters struct {
	Source string
	Status string
	Limit  string
	Days   string
}

// Get a list of all of the event categories used by the EONET API
func GetCategories() ([]byte, error) {
	responseData, err := sendRequestToEonetApi(baseCategoriesUrl)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

// Get a list of all the events under a specific category. if category is == ""
// then the behavior will be the same as calling GetCategories()
func GetEventsByCategory(category string, query CategoriesQueryParameters) ([]byte, error) {
	url := createCategoriesApiUrl(category, query)

	responseData, err := sendRequestToEonetApi(url.String())
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func createCategoriesApiUrl(category string, query CategoriesQueryParameters) url.URL {
	if category == "" {
		return url.URL{Scheme: "https", Host: "eonet.sci.gsfc.nasa.gov", Path: "/api/v3/categories"}
	}

	u := url.URL{
		Scheme: "https",
		Host:   "eonet.sci.gsfc.nasa.gov",
		Path:   "/api/v3/categories/" + category,
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
	return u
}
