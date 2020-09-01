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

func TestGetEventsWithLimit(t *testing.T) {
  mockCtrl := gomock.NewController(t)

  mockHTTPClient := mocks.NewMockHTTPClient(mockCtrl)
  client = mockHTTPClient

  request, _ := http.NewRequest("GET", "https://eonet.sci.gsfc.nasa.gov/api/v3/events?limit=1", nil)

  response := &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte("")))}

  mockHTTPClient.EXPECT().Do(gomock.Eq(request)).Return(response, nil).Times(1)

  GetEvents(EventsQueryParameters{Limit: 1})

  client = &http.Client{Timeout: 5 * time.Second}
}

func TestGetEventsForCorrectUrl(t *testing.T) {
  mockCtrl := gomock.NewController(t)

  mockHTTPClient := mocks.NewMockHTTPClient(mockCtrl)
  client = mockHTTPClient

  url := "https://eonet.sci.gsfc.nasa.gov/api/v3/events?bbox=-129.02%2C50.73%2C-58.71%2C12.89&days=20&end=2019-01-31&magID=mag_kts&magMax=20&magMin=1.50&start=2019-01-01"

  request, err := http.NewRequest("GET", url, nil)

  if err != nil {
    t.Error(err)
  }

  response := &http.Response{Body: ioutil.NopCloser(bytes.NewReader([]byte("")))}

  mockHTTPClient.EXPECT().Do(gomock.Eq(request)).Return(response, nil).Times(1)

  query := EventsQueryParameters {
    Days: 20,
    Start: "2019-01-01",
    End: "2019-01-31",
    MagID: "mag_kts",
    MagMin: "1.50",
    MagMax: "20",
    Bbox: "-129.02,50.73,-58.71,12.89",
  }

  GetEvents(query)

  client = &http.Client{Timeout: 5 * time.Second}
}
