package internal

import (
	"fmt"
	"os"

	"github.com/DoeMoor/pokedexcli/internal/locations_map"
)

func commandExit() error {
	fmt.Println("Exiting Pokedex")
	ClearTerminal()
	os.Exit(0)
	return nil
}

func commandHelp() error {
	commandsList := GetCliCommandsList()
	fmt.Println("Welcome to the Pokedex!\n\n Usage:")
	for key, command := range commandsList {
		fmt.Printf("  %s: %s\n", key, command.description)
	}
	fmt.Println("")
	return nil
}

func commandMap() error {
	locations_map.NextLocations()
	return nil
}

func commandMapb() error {
	locations_map.PreviousLocation()
	return nil
}
