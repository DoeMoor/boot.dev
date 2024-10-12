package command

import (
	"fmt"
)

func commandHelp(userInputList []string) error {

	commandsList := GetCliCommandsList()

	fmt.Println("Welcome to the Pokedex!"," Usage:")
	for _, cmd := range commandsList {
		fmt.Printf("  %s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println("")
	return nil
}