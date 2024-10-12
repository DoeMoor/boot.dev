package command

import(
	"fmt"
	"os"
	"github.com/DoeMoor/pokedexcli/internal/utility"
)

func commandExit(userInputList []string) error {
	fmt.Println("Exiting Pokedex")
	utility.ClearTerminal()
	os.Exit(0)
	return nil
}