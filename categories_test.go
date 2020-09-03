package goeonet

import (
  "bytes"
  "io/ioutil"
  "net/http"
  "testing"

  "github.com/golang/mock/gomock"
  "github.com/jfitzg7/goeonet/mocks"
  "github.com/onsi/gomega"
)

const mockCategoriesJsonData = `{
	"title": "EONET Event Categories",
	"description": "List of all the available event categories in the EONET system",
	"link": "https://eonet.sci.gsfc.nasa.gov/api/v3/categories",
	"categories": [
		{
			"id": "drought",
			"title": "Drought",
			"link": "https://eonet.sci.gsfc.nasa.gov/api/v3/categories/drought",
			"description": "Long lasting absence of precipitation affecting agriculture and livestock, and the overall availability of food and water.",
			"layers": "https://eonet.sci.gsfc.nasa.gov/api/v3/layers/drought"
		}
  ]
}`

func TestGetCategories(t *testing.T) {
  mockCtrl := gomock.NewController(t)

  mockHTTPClient := mocks.NewMockHTTPClient(mockCtrl)
  client = mockHTTPClient

  url := "https://eonet.sci.gsfc.nasa.gov/api/v3/categories"

  request, _ := http.NewRequest("GET", url, nil)

  response := &http.Response{Body: ioutil.NopCloser(bytes.NewReader([]byte(mockCategoriesJsonData)))}

  mockHTTPClient.EXPECT().Do(gomock.Eq(request)).Return(response, nil).Times(1)

  jsonData, err := GetCategories()

  if err != nil {
    t.Error(err)
  }

  g := gomega.NewGomegaWithT(t)
  g.Expect(string(jsonData)).To(gomega.MatchJSON(mockCategoriesJsonData))
}

func TestGetEventsByCategory(t *testing.T) {
  mockCtrl := gomock.NewController(t)

  mockHTTPClient := mocks.NewMockHTTPClient(mockCtrl)
  client = mockHTTPClient

  url := "https://eonet.sci.gsfc.nasa.gov/api/v3/categories/wildfires?days=30&limit=1&source=InciWeb&status=open"

  request, _ := http.NewRequest("GET", url, nil)

  response := &http.Response{Body: ioutil.NopCloser(bytes.NewReader([]byte(mockCategoriesJsonData)))}

  mockHTTPClient.EXPECT().Do(gomock.Eq(request)).Return(response, nil).Times(1)

  query := CategoriesQueryParameters{
    Category: "wildfires",
    Source: "InciWeb",
    Status: "open",
    Limit: "1",
    Days: "30",
  }

  jsonData, err := GetEventsByCategory(query)

  if err != nil {
    t.Error(err)
  }

  g := gomega.NewGomegaWithT(t)
  g.Expect(string(jsonData)).To(gomega.MatchJSON(mockCategoriesJsonData))
}
