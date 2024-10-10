package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/DoeMoor/pokedexcli/internal/command"
	"github.com/DoeMoor/pokedexcli/internal/utility"
)

func REPL() {

	utility.ClearTerminal()

	commandsList := command.GetCliCommandsList()
	
	for {
		fmt.Print("Pokedex > ")

		stdScanner := bufio.NewScanner(os.Stdin)
		stdScanner.Scan()
		userInput := stdScanner.Text()

		// check if the user input is empty
		if userInput == "" {
			commandsList["help"].Callback()
			continue
		}

		// check if the command is in the list
		if _, ok := commandsList[userInput]; !ok {
			fmt.Println(" not existed command:  ", userInput)
			commandsList["help"].Callback()
			continue
		}

		
		// execute the command
		commandsList[userInput].Callback()

	}

}
