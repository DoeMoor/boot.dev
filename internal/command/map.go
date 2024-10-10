package command

import (
	"fmt"

	"github.com/DoeMoor/pokedexcli/internal/client"
	"github.com/DoeMoor/pokedexcli/internal/json_scheme"
)

func NextLocations() error {

	url, err := callLocalEndpoint("next")
	if err != nil {
		fmt.Println(err)
		return err
	}

	var loc json_scheme.Locations

	err = client.ApiCall(url, &loc)

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

	err = client.ApiCall(url, &loc)

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

	if cfg.CurrentURL == "" {
		cfg.CurrentURL = fmt.Sprintf("%v%v?offset=%v&limit=%v",
			cfg.BaseURL,
			locationsEndpoint,
			cfg.Offset, cfg.Limit)

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

func printLocations(loc json_scheme.Locations) {
	fmt.Println("Locations:")
	for _, location := range loc.Results {
		fmt.Printf(" %v\n", location.Name)
	}
	fmt.Println("")
}
