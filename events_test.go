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

func TestGetEventsWithSource(t *testing.T) {
  mockCtrl := gomock.NewController(t)

  mockHTTPClient := mocks.NewMockHTTPClient(mockCtrl)
  client = mockHTTPClient

  request, _ := http.NewRequest("GET", "https://eonet.sci.gsfc.nasa.gov/api/v3/events?limit=1&source=InciWeb", nil)

  response := &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte("")))}

  mockHTTPClient.EXPECT().Do(gomock.Eq(request)).Return(response, nil).Times(1)

  GetEvents(EventsQueryParameters{Source: "InciWeb", Limit: 1})

  client = &http.Client{Timeout: 5 * time.Second}
}
