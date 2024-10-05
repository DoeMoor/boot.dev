package locations_map

import (
	"fmt"
	"net/http"
)

type URLconfig struct {
	currentURL string
	Next       string
	Previous   string
	limit      int
	offset     int
}

var Config *URLconfig

func GetURLconfig() *URLconfig {
	// if RequestConfig is nil, set the default values
	if Config == nil {
		Config = &URLconfig{
			currentURL: "https://pokeapi.co/api/v2/location?offset=0&limit=5",
			Next:       "https://pokeapi.co/api/v2/location?offset=0&limit=5",
			offset:     0,
			limit:      5,
		}
	}
	return Config
}

func (conf *URLconfig) SetURL(loc locations, url string) {
	conf.currentURL = url

	if loc.Next == nil {
		conf.Next = ""
	} else {
		conf.Next = *loc.Next
	}

	if loc.Previous == nil {
		conf.Previous = ""
	} else {
		conf.Previous = *loc.Previous
	}
}

func (URLconf *URLconfig) SetQueryParam(limit, offset int) {
	URLconf.limit = limit
	URLconf.offset = offset
}

func NextLocations() error {
	url, err := getURL("next")

	if err != nil {
		fmt.Println(err)
		return err
	}

	loc, err := getLocations(url)
	if err != nil {
		fmt.Println(err)
		return err
	}

	printLocations(loc)
	GetURLconfig().SetURL(loc, url)
	GetCache().write(url, loc)

	return nil
}

func PreviousLocation() error {
	url, err := getURL("previous")
	if err != nil {
		fmt.Println(err)
		return err
	}

	loc, err := getLocations(url)
	if err != nil {
		fmt.Println(err)
		return err
	}

	printLocations(loc)
	GetURLconfig().SetURL(loc, url)
	GetCache().write(url, loc)
	return nil
}

func getURL(direction string) (url string, err error) {
	cfg := GetURLconfig()
	switch direction {
	case "next":
		if cfg.Next == "" {
			return cfg.currentURL, fmt.Errorf("no next location")
		}
		return cfg.Next, nil

	case "previous":
		if cfg.Previous == "" {
			return cfg.currentURL, fmt.Errorf("no previous location")
		}
		return cfg.Previous, nil

	default:
		return "", fmt.Errorf("invalid direction")
	}
}

func getLocations(url string) (locations, error) {
	var loc locations

	loc, ok := GetCache().read(url)
	if ok {
		fmt.Println("Cache hit")
		return loc, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locations{}, err
	}

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return locations{}, err
	}
	defer resp.Body.Close()

	err = loc.newLocations(resp)
	if err != nil {
		return locations{}, err
	}

	return loc, nil
}

func printLocations(loc locations) {
	fmt.Println("Locations:")
	for _, location := range loc.Results {
		fmt.Printf(" %v\n", location.Name)
	}
	fmt.Println("")
}
