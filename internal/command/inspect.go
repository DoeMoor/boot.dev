package command

import (
	"fmt"

	"github.com/DoeMoor/pokedexcli/internal/db"
)

func inspect(userInput []string) error {
	if len(userInput) < 2 {
		fmt.Println("Please provide a Pokemon name to inspect")
		return nil
	}

	pokemonName := userInput[1]
	pokemon, err := db.GetUserInventory().GetPokemonFromPokedex(pokemonName)
	if err != nil {
		return err
	}

	fmt.Println("Inspecting", pokemon.Name, "...")
	fmt.Printf(`Name: %v
Height: %v
Weight: %v
`, pokemon.Name,
		pokemon.Height,
		pokemon.Weight)
	
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%v:%v\n",stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Println(" -", t.Type.Name)
	}
	return nil
}
