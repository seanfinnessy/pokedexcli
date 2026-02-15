package main

type cliCommand struct {
	name string
	description string
	callback func(*AppState, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Display all available locations. Call again to go to next page.",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Display all available locations. Call to move back a page.",
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "See a list of all Pokemon inside a location by passing the name of the location as an argument.",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Catch a pokemon by passing the name of the Pokemon as an argument.",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Inspect a Pokemon in your Pokedex.",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "View your Pokedex.",
			callback: commandPokedex,
		},
	}
}
