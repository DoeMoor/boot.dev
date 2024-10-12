package command

import (
	"fmt"

	"github.com/DoeMoor/pokedexcli/internal/db"
)

func pokedex(userInput []string) error {
	pokedex, err := db.GetUserInventory().GetPokemonsListFromPokedex()
	if err != nil {
		return err
	}
	
	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokedex {
		fmt.Println(" - "+pokemon.Name)
	}
	return nil
}