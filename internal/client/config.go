package client

type clientConfig struct {
	BaseURL    string
	CurrentURL string
	Next       string
	Previous   string
	Limit      int
	Offset     int
}

var config *clientConfig

// GetClientConfig returns the singleton current client configuration
func GetClientConfig() *clientConfig {
	// if RequestConfig is nil, set the default values
	if config == nil {
		config = &clientConfig{
			BaseURL: "https://pokeapi.co/api/v2",
			Limit:   5,
			Offset:  0,
		}
	}
	return config
}

// SetURL sets the current URL and the next and previous URLs
func (conf *clientConfig) SetURL(current string, next *string, previous *string) {
	conf.CurrentURL = current

	if next == nil {
		conf.Next = ""
	} else {
		conf.Next = *next
	}

	if previous == nil {
		conf.Previous = ""
	} else {
		conf.Previous = *previous
	}
}

func (URLconf *clientConfig) SetQueryParam(limit, offset int) {
	URLconf.Limit = limit
	URLconf.Offset = offset
}
