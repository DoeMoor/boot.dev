package command

import (
	"fmt"

	"github.com/DoeMoor/pokedexcli/internal/client"
	"github.com/DoeMoor/pokedexcli/internal/json_scheme"
)



// type URLconfig struct {
// 	currentURL string
// 	Next       string
// 	Previous   string
// 	limit      int
// 	offset     int
//  endpoint   map{string}[]EndpointParam
// }

// var Config *URLconfig

// func GetURLconfig() *URLconfig {
// 	// if RequestConfig is nil, set the default values
// 	if Config == nil {
// 		Config = &URLconfig{
// 			currentURL: "https://pokeapi.co/api/v2/location?offset=0&limit=5",
// 			Next:       "https://pokeapi.co/api/v2/location?offset=0&limit=5",
// 			offset:     0,
// 			limit:      5,
// 		}
// 	}
// 	return Config
// }

// // SetURL sets the current URL and the next and previous URLs
// // based on the location struct
// func (conf *URLconfig) SetURL(loc sch.Locations , url string) {
// 	conf.currentURL = url

// 	if loc.Next == nil {
// 		conf.Next = ""
// 	} else {
// 		conf.Next = *loc.Next
// 	}

// 	if loc.Previous == nil {
// 		conf.Previous = ""
// 	} else {
// 		conf.Previous = *loc.Previous
// 	}
// }

// func (URLconf *URLconfig) SetQueryParam(limit, offset int) {
// 	URLconf.limit = limit
// 	URLconf.offset = offset
// }

func NextLocations() error {

	url, err := callLocalEndpoint("next")
	if err != nil {
		fmt.Println(err)
		return err
	}

	var loc json_scheme.Locations

	err = client.ApiCall(url,&loc)

	if err != nil {
		fmt.Println(err)
		return err
	}

	printLocations(loc)
	client.GetClientConfig().SetURL(url, loc.Next, loc.Previous)
	return nil
}

func PreviousLocation() error {
	url, err := callLocalEndpoint("previous")
	if err != nil {
		fmt.Println(err)
		return err
	}

	var loc json_scheme.Locations

	err = client.ApiCall(url,&loc)
	
	if err != nil {
		fmt.Println(err)
		return err
	}

	printLocations(loc)
	client.GetClientConfig().SetURL(url, loc.Next, loc.Previous)
	return nil
}

func callLocalEndpoint(direction string) (string, error) {
	cfg := client.GetClientConfig()
	locationsEndpoint := "/location-area"

	if cfg.CurrentURL == ""{
		cfg.CurrentURL = fmt.Sprintf("%v%v?offset=%v&limit=%v",
		cfg.BaseURL,
		locationsEndpoint,
		cfg.Offset,cfg.Limit)

		return cfg.CurrentURL, nil
	}

	if direction == "next" {
		if cfg.Next == "" {
			return cfg.CurrentURL, fmt.Errorf("no next location")
		}
		return cfg.Next, nil
	}

	if direction == "previous" {
		if cfg.Previous == "" {
			return cfg.CurrentURL, fmt.Errorf("no previous location")
		}
		return cfg.Previous, nil
	}
	return cfg.CurrentURL, fmt.Errorf("something wrong in callLocalEndpoint()")
}

//  func getLocations(url string) (sch.Locations, error) {
// 	var loc sch.Locations

// 	cached, ok := cache.GetCache().Read(url)
// 	if ok {
// 		fmt.Println("Cache hit")
// 		loc.newLocations(cached)
// 		return loc, nil
// 	}

// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return locations{}, err
// 	}

// 	httpClient := http.Client{}
// 	resp, err := httpClient.Do(req)
// 	if err != nil {
// 		return locations{}, err
// 	}

// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("cache write error:", err)
// 	} else {
// 		pokecache.GetCache().Write(url, body)
// 	}

// 	err = loc.newLocations(body)
// 	if err != nil {
// 		return locations{}, err
// 	}
// 	return loc, nil
// }

func printLocations(loc json_scheme.Locations) {
	fmt.Println("Locations:")
	for _, location := range loc.Results {
		fmt.Printf(" %v\n", location.Name)
	}
	fmt.Println("")
}
