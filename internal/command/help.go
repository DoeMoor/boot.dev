package command

import (
	"fmt"
)

func commandHelp() error {

	commandsList := GetCliCommandsList()

	fmt.Println("Welcome to the Pokedex!\n\n Usage:")
	for key, command := range commandsList {
		fmt.Printf("  %s: %s\n", key, command.Description)
	}
	fmt.Println("")
	return nil
}