package command

import (
	"fmt"
	"github.com/DoeMoor/pokedexcli/internal/client"
	"github.com/DoeMoor/pokedexcli/internal/endpoint_scheme"
)

func explore(userInputList []string) error {
	if len(userInputList) < 2 {
		fmt.Printf("%v command does not take any arguments \n", userInputList[0])
		return nil
	}
	cnf := client.GetClientConfig()
	locationEndpointURL := cnf.BaseURL + "/location-area/" + userInputList[1] + "/"
	var loc endpoint_scheme.LocationArea
	err := client.ApiCall(locationEndpointURL, &loc)
	if err != nil {
		return err
	}
	printDiscoveredPokemonName(loc)
	return nil
}

func printDiscoveredPokemonName(loc endpoint_scheme.LocationArea) {
	fmt.Println("Exploring ", loc.Name, "...")
	fmt.Println("Found Pokemon:")
	for _, location := range loc.PokemonEncounters {
		fmt.Println(" - ", location.Pokemon.Name)
	}
}
