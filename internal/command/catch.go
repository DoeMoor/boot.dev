package command

import (
	"fmt"
	"math/rand"

	"github.com/DoeMoor/pokedexcli/internal/client"
	"github.com/DoeMoor/pokedexcli/internal/db"
	"github.com/DoeMoor/pokedexcli/internal/endpoint_scheme"
)

func catch(userInput []string) error {
	if len(userInput) < 2 {
		fmt.Println("Please provide a Pokemon name to catch")
		return nil
	}
	pokemonName := userInput[1]
	cnf := client.GetClientConfig()
	endpointURL := cnf.BaseURL + "/pokemon/" + pokemonName + "/"
	var pokemon endpoint_scheme.Pokemon

	err := client.ApiCall(endpointURL, &pokemon)
	if err != nil {
		return err
	}

	throwPokeball(pokemon)

	return nil
}

func throwPokeball(pokemon endpoint_scheme.Pokemon) {
	fmt.Println("Throwing a pokeball at", pokemon.Name, "...")

	if isCaught(pokemon.BaseExperience) {
		fmt.Println(pokemon.Name, "was caught!")
		db.GetUserInventory().AddPokemon(pokemon)
	} else {
		fmt.Println(pokemon.Name, "escaped!")
	}
}

func isCaught(experience int64) bool {
	if rand.Intn(100) < int(experience) {
		return true
	}
	return false
}