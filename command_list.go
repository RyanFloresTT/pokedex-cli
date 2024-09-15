package main

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, input string) error
}

var commands map[string]cliCommand

func initCommands() {
	Pokedex = make(map[string]Pokemon)

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
		"catch": {
			name:        "catch {pokemon}",
			description: "Attempts to catch the specified Pokemon.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect {pokemon}",
			description: "Inspects a pokemon from your Pokedex.",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists the pokedex in your Pokedex.",
			callback:    commandPokedex,
		},
	}
}
