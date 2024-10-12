package utility

import (
	"strings"
)

func cleanUserCommandInput(userInput string) (string) {
	userInput = strings.ReplaceAll(userInput, "/", "")
	userInput = strings.ReplaceAll(userInput, "\\", "")
	userInput = strings.ToLower(userInput)
	return userInput
}

func ListOfUserInput(userInput string) ([]string) {
	clearInput := cleanUserCommandInput(userInput)
	input := strings.Split(clearInput, " ")
	return input
}