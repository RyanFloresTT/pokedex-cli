package main

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, input string) error
}

var commands map[string]cliCommand

func initCommands() {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Shows the next 20 map locations.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous 20 map locations.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore {area}",
			description: "Shows Pokemon encounters in the given area.",
			callback:    commandExplore,
		},
	}
}
