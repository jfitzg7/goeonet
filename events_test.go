package goeonet

import (
  "bytes"
  "io/ioutil"
  "testing"
  "net/http"
  "time"

  "github.com/golang/mock/gomock"
  "github.com/jfitzg7/goeonet/mocks"
)

const mockEventsJsonExample = `{
  "title": "EONET Events",
  "description": "Natural events from EONET.",
  "link": "https://eonet.sci.gsfc.nasa.gov/api/v3/events",
  "events": [
    {
      "id": "EONET_4954",
      "title": "Deep Creek Fire",
      "description": null,
      "link": "https://eonet.sci.gsfc.nasa.gov/api/v3/events/EONET_4954",
      "closed": null,
      "categories": [
        {
          "id": "wildfires",
          "title": "Wildfires"
        }
      ],
      "sources": [
        {
          "id": "InciWeb",
          "url": "http://inciweb.nwcg.gov/incident/7112/"
        }
      ],
      "geometry": [
        {
          "magnitudeValue": null,
          "magnitudeUnit": null,
          "date": "2020-08-30T08:43:00Z",
          "type": "Point",
          "coordinates": [ -99.150000000000006, 32.686999999999998 ]
        }
      ]
    }
  ]
}`

func TestGetEventsWithSource(t *testing.T) {
  mockCtrl := gomock.NewController(t)

  mockHTTPClient := mocks.NewMockHTTPClient(mockCtrl)
  client = mockHTTPClient

  url := "https://eonet.sci.gsfc.nasa.gov/api/v3/events?end=2020-08-30&limit=1&source=InciWeb&start=2020-08-30&status=open"

  request, _ := http.NewRequest("GET", url, nil)

  response := &http.Response{Body: ioutil.NopCloser(bytes.NewReader([]byte(mockEventsJsonExample)))}

  mockHTTPClient.EXPECT().Do(gomock.Eq(request)).Return(response, nil).Times(1)

  query := EventsQueryParameters{
    Source: "InciWeb",
    Status: "open",
    Limit: 1,
    Start: "2020-08-30",
    End: "2020-08-30",
  }

  _, err := GetEvents(query)

  if err != nil {
    t.Error(err)
  }

  client = &http.Client{Timeout: 5 * time.Second}
}

func TestGetEventsForCorrectUrl(t *testing.T) {
  mockCtrl := gomock.NewController(t)

  mockHTTPClient := mocks.NewMockHTTPClient(mockCtrl)
  client = mockHTTPClient

  url := "https://eonet.sci.gsfc.nasa.gov/api/v3/events?bbox=-129.02%2C50.73%2C-58.71%2C12.89&days=20&magID=mag_kts&magMax=20&magMin=1.50"

  request, _ := http.NewRequest("GET", url, nil)

  response := &http.Response{Body: ioutil.NopCloser(bytes.NewReader([]byte("")))}

  mockHTTPClient.EXPECT().Do(gomock.Eq(request)).Return(response, nil).Times(1)

  query := EventsQueryParameters {
    Days: 20,
    MagID: "mag_kts",
    MagMin: "1.50",
    MagMax: "20",
    Bbox: "-129.02,50.73,-58.71,12.89",
  }

  _, err := GetEvents(query)

  if err != nil {
    t.Error(err)
  }

  client = &http.Client{Timeout: 5 * time.Second}
}
