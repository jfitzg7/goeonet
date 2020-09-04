package goeonet

const baseMagnitudesUrl = "https://eonet.sci.gsfc.nasa.gov/api/v3/magnitudes"

// GetMagnitudes gets a list of all the units used for measuring
// the magnitudes of specific events
func GetMagnitudes() ([]byte, error) {
	responseData, err := sendRequestToEonetApi(baseMagnitudesUrl)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
