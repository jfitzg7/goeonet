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

const mockLayersJsonData = `{
  "title": "EONET Web Service Layers",
  "description": "List of web service layers in the EONET system",
  "link": "https://eonet.sci.gsfc.nasa.gov/api/v3/layers",
  "categories": [
    {
      "layers": [
        {
          "name": "AIRS_CO_Total_Column_Day",
          "serviceUrl": "https://gibs.earthdata.nasa.gov/wmts/epsg4326/best/wmts.cgi",
          "serviceTypeId": "WMTS_1_0_0",
          "parameters": [
            {
              "TILEMATRIXSET": "2km",
              "FORMAT": "image/png"
            }
          ]
        }
      ]
    }
  ]
}`

func TestGetLayers(t *testing.T) {
  mockCtrl := gomock.NewController(t)

  mockHTTPClient := mocks.NewMockHTTPClient(mockCtrl)
  client = mockHTTPClient

  url := "https://eonet.sci.gsfc.nasa.gov/api/v3/layers"

  request, _ := http.NewRequest("GET", url, nil)

  response := &http.Response{Body: ioutil.NopCloser(bytes.NewReader([]byte(mockLayersJsonData)))}

  mockHTTPClient.EXPECT().Do(gomock.Eq(request)).Return(response, nil).Times(1)

  jsonData, err := GetLayers()

  if err != nil {
    t.Error(err)
  }

  g := gomega.NewGomegaWithT(t)
  g.Expect(string(jsonData)).To(gomega.MatchJSON(mockLayersJsonData))
}

func TestGetLayersByCategoryID(t *testing.T) {
  mockCtrl := gomock.NewController(t)

  mockHTTPClient := mocks.NewMockHTTPClient(mockCtrl)
  client = mockHTTPClient

  url := "https://eonet.sci.gsfc.nasa.gov/api/v3/layers/wildfires"

  request, _ := http.NewRequest("GET", url, nil)

  response := &http.Response{Body: ioutil.NopCloser(bytes.NewReader([]byte(mockLayersJsonData)))}

  mockHTTPClient.EXPECT().Do(gomock.Eq(request)).Return(response, nil).Times(1)

  jsonData, err := GetLayersByCategoryID("wildfires")

  if err != nil {
    t.Error(err)
  }

  g := gomega.NewGomegaWithT(t)
  g.Expect(string(jsonData)).To(gomega.MatchJSON(mockLayersJsonData))
}
