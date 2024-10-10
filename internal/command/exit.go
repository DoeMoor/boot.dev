package command

import(
	"fmt"
	"os"
	"github.com/DoeMoor/pokedexcli/internal/utility"
)

func commandExit() error {
	fmt.Println("Exiting Pokedex")
	utility.ClearTerminal()
	os.Exit(0)
	return nil
}