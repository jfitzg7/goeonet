package goeonet

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jfitzg7/goeonet/mocks"
	"github.com/onsi/gomega"
)

const mockMagnitudesJsonData = `{
	"title": "EONET Event Magnitudes",
	"description": "List of all the available event magnitudes in the EONET system",
	"link": "https://eonet.sci.gsfc.nasa.gov/api/v3/magnitudes",
	"magnitudes": [
		{
			"id": "mag_kts",
			"name": "Avg Max Windspeed (kts)",
			"unit": "kts",
			"description": "Average Max Sustained Winds reported for severe storms.",
			"link": "https://eonet.sci.gsfc.nasa.gov/api/v3/events?magID=mag_kts"
		},

			{
			"id": "mms",
			"name": "Moment Magnitude Scale",
			"unit": "Mw",
			"description": "Moment magnitude scale (MMS) denoted as Mw : measure of an earthquake.",
			"link": "https://eonet.sci.gsfc.nasa.gov/api/v3/events?magID=mms"
		},

			{
			"id": "sq_NM",
			"name": "Area (Nautical Miles)",
			"unit": "NM^2",
			"description": "Nautical miles squared used to measure area of icebergs.",
			"link": "https://eonet.sci.gsfc.nasa.gov/api/v3/events?magID=sq_NM"
		}

		]
}`

func TestGetMagnitudes(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	mockHTTPClient := mocks.NewMockHTTPClient(mockCtrl)
	client = mockHTTPClient

	url := "https://eonet.sci.gsfc.nasa.gov/api/v3/magnitudes"

	request, _ := http.NewRequest("GET", url, nil)

	response := &http.Response{Body: ioutil.NopCloser(bytes.NewReader([]byte(mockMagnitudesJsonData)))}

	mockHTTPClient.EXPECT().Do(gomock.Eq(request)).Return(response, nil).Times(1)

	jsonData, err := GetMagnitudes()

	if err != nil {
		t.Error(err)
	}

	g := gomega.NewGomegaWithT(t)
	g.Expect(string(jsonData)).To(gomega.MatchJSON(mockMagnitudesJsonData))
}

func TestGetGetMagnitudesError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	mockHTTPClient := mocks.NewMockHTTPClient(mockCtrl)
	client = mockHTTPClient

	url := "https://eonet.sci.gsfc.nasa.gov/api/v3/magnitudes"

	request, _ := http.NewRequest("GET", url, nil)

	mockHTTPClient.EXPECT().Do(gomock.Eq(request)).Return(nil, errors.New("mock error")).Times(1)

	_, err := GetMagnitudes()

	if err == nil {
		t.Error(errors.New("An error should have occurred"))
	}
}
