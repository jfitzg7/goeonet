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

const mockSourcesJsonData = `{
	"title": "EONET Event Sources",
	"description": "List of all the available event sources in the EONET system",
	"link": "https://eonet.sci.gsfc.nasa.gov/api/v3/sources",
	"sources": [
    {
			"id": "AVO",
			"title": "Alaska Volcano Observatory",
			"source": "https://www.avo.alaska.edu/",
			"link": "https://eonet.sci.gsfc.nasa.gov/api/v3/events?source=AVO"
		}
  ]
}`

func TestGetSources(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	mockHTTPClient := mocks.NewMockHTTPClient(mockCtrl)
	client = mockHTTPClient

	url := "https://eonet.sci.gsfc.nasa.gov/api/v3/sources"

	request, _ := http.NewRequest("GET", url, nil)

	response := &http.Response{Body: ioutil.NopCloser(bytes.NewReader([]byte(mockSourcesJsonData)))}

	mockHTTPClient.EXPECT().Do(gomock.Eq(request)).Return(response, nil).Times(1)

	jsonData, err := GetSources()

	if err != nil {
		t.Error(err)
	}

	g := gomega.NewGomegaWithT(t)
	g.Expect(string(jsonData)).To(gomega.MatchJSON(mockSourcesJsonData))
}
