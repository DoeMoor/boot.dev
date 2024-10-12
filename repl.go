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
		userInputList := utility.ListOfUserInput(stdScanner.Text())

		// check if the user input is empty
		if userInputList[0] == "" {
			commandsList["help"].Callback(userInputList)
			continue
		}

		// check if the command is in the list
		if _, ok := commandsList[userInputList[0]]; !ok {
			fmt.Println(" not existed command:  ", userInputList)
			commandsList["help"].Callback(userInputList)
			continue
		}

		// execute the command
		err := commandsList[userInputList[0]].Callback(userInputList)
		if err != nil {
			fmt.Println("error: ", err)
		}

	}

}
