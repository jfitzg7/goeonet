package goeonet

type Source struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Source string `json:"source"`
	Link   string `json:"link"`
}

type Sources struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Link        string   `json:"link"`
	Sources     []Source `json:"sources"`
}

func GetSources() (*Collection, error) {
	collection, err := queryEonetApi(baseSourcesUrl)
	if err != nil {
		return nil, err
	}

	return collection, nil
}
