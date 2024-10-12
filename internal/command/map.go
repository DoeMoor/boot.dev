package command

import (
	"fmt"

	"github.com/DoeMoor/pokedexcli/internal/client"
	"github.com/DoeMoor/pokedexcli/internal/endpoint_scheme"
)

func nextLocations(userInputList []string) error {

	url, err := callLocalEndpoint("next")
	if err != nil {
		return err
	}

	var loc endpoint_scheme.Locations
	err = client.ApiCall(url, &loc)
	if err != nil {
		return err
	}

	printLocations(loc)
	client.GetClientConfig().SetURL(url, loc.Next, loc.Previous)
	return nil
}

func previousLocation(userInputList []string) error {
	url, err := callLocalEndpoint("previous")
	if err != nil {
		return err
	}

	var loc endpoint_scheme.Locations
	err = client.ApiCall(url, &loc)
	if err != nil {
		return err
	}

	printLocations(loc)
	client.GetClientConfig().SetURL(url, loc.Next, loc.Previous)
	return nil
}

func callLocalEndpoint(direction string) (string, error) {
	cfg := client.GetClientConfig()
	locationsEndpoint := "/location-area"

	if cfg.CurrentURL == "" {
		cfg.CurrentURL = fmt.Sprintf("%v%v?offset=%v&limit=%v",
			cfg.BaseURL,
			locationsEndpoint,
			cfg.Offset, cfg.Limit)

		return cfg.CurrentURL, nil
	}

	if direction == "next" {
		if cfg.Next == "" {
			fmt.Println("no next location")
			return cfg.CurrentURL, nil
		}
		return cfg.Next, nil
	}

	if direction == "previous" {
		if cfg.Previous == "" {
			fmt.Println("no previous location")
			return cfg.CurrentURL, nil
		}
		return cfg.Previous, nil
	}
	return cfg.CurrentURL, fmt.Errorf("something wrong in callLocalEndpoint()")
}

func printLocations(loc endpoint_scheme.Locations) {
	fmt.Println("Locations:")
	for _, location := range loc.Results {
		fmt.Printf(" %v\n", location.Name)
	}
	fmt.Println("")
}
