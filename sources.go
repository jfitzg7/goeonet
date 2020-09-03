package goeonet

const	baseSourcesUrl = "https://eonet.sci.gsfc.nasa.gov/api/v3/sources"

func GetSources() ([]byte, error) {
	responseData, err := sendRequestToEonetApi(baseSourcesUrl)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
