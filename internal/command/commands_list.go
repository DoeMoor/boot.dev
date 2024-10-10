package command



type cliCommands struct {
	Name string
	Description string
	Callback func() error
}

//command list
// GetCliCommandsList returns a map of commands
	//	"help": {
	// 	description: "Prints the help menu",
	// 	Callback: commandHelp,
	// "exit","q": {
	// 	description: "Exits the Pokedex",
	// 	Callback: commandExit,
	// "map": {
	// 	description: "Shows the map of the current location",
	// 	Callback: commandMap,
	// 	"mapb":
	// 	description: "Shows the map of the previous location",
	// 	Callback: commandMapb,
func GetCliCommandsList() map[string]cliCommands {
	return map[string]cliCommands{
		"help": {
			Name: "help",
			Description: "Prints the help menu",
			Callback: commandHelp,
		},
		"exit": {
			Name: "exit",
			Description: "Exits the Pokedex",
			Callback: commandExit,
		},
		"q": {
			Name: "q",
			Description: "alias - Exits the Pokedex",
			Callback: commandExit,
		},
		"map": {
			Name: "map",
			Description: "Shows the map of the current location",
			Callback: NextLocations,
		},
		"mapb": {
			Name: "mapb",
			Description: "Shows the map of the previous location",
			Callback: PreviousLocation,
		},
	}
}



