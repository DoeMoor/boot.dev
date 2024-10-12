package command

import "github.com/DoeMoor/pokedexcli/internal/utility"

func clear(userInputList []string) error {
	utility.ClearTerminal()
	return nil
}