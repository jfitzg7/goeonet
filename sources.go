package goeonet

const baseSourcesUrl = "https://eonet.sci.gsfc.nasa.gov/api/v3/sources"

// Get a list of all the sources used by the EONET API
func GetSources() ([]byte, error) {
	responseData, err := sendRequestToEonetApi(baseSourcesUrl)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
