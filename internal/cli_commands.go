package internal

type cliCommand struct {
	name string
	description string
	Callback func() error
}

//command list
// GetCliCommandsList returns a map of commands
	//	"help": {
	// 	description: "Prints the help menu",
	// 	Callback: commandHelp,
	// "exit": {
	// 	description: "Exits the Pokedex",
	// 	Callback: commandExit,
	// "map": {
	// 	description: "Shows the map of the current location",
	// 	Callback: commandMap,
	// 	"mapb":
	// 	description: "Shows the map of the previous location",
	// 	Callback: commandMapb,
func GetCliCommandsList() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Prints the help menu",
			Callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exits the Pokedex",
			Callback: commandExit,
		},
		"q": {
			name: "q",
			description: "alias - Exits the Pokedex",
			Callback: commandExit,
		},
		"map": {
			name: "map",
			description: "Shows the map of the current location",
			Callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Shows the map of the previous location",
			Callback: commandMapb,
		},
	}
}



