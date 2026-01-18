package main

import "github.com/seanfinnessy/pokedexcli/internal/pokeapi"

type cliCommand struct {
	name string
	description string
	callback func(*pokeapi.LocationAreaResObject) error
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
	}
}
