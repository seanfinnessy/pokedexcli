package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	pokeapi "github.com/seanfinnessy/pokedexcli/internal/pokeapi"
)

var config pokeapi.LocationAreaResObject

func startRepl() {
	// Create a new scanner
	scanner := bufio.NewScanner(os.Stdin)
	
	// REPL
	for {
		fmt.Print("Pokedex > ")
		// Scan for input
		scanner.Scan()
		// Apply input to variable
		input := scanner.Text()
		// Clean input, grab first word, set command and show user
		cleanedInput := cleanInput(input)
		command := cleanedInput[0]
		
		// Verify command
		checkCommand(command)
	}
}

func checkCommand(commandString string) {
	command, ok := getCommands()[commandString]
	if !ok {
		fmt.Println("Unknown command.")
	} else {
		// callback function, pass addr to config
		command.callback(&config)
	}
}

func commandExit(config *pokeapi.LocationAreaResObject) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *pokeapi.LocationAreaResObject) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println("")

	for _, value  := range getCommands() {
		helpMsg := fmt.Sprintf("%s: %s", value.name, value.description)
		fmt.Println(helpMsg)
	}
	return nil
}

func commandMap(config *pokeapi.LocationAreaResObject) error {
	var url string
	
	// If next is nil (aka first time using map command). We set it to the first page.
	if config.Next == nil {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	// If not nil, we set the url to search for Next Page.
	if config.Next != nil {	
		url = *config.Next	
	}

	// Call API, pass in the URL to be searched
	err := pokeapi.GetLocationAreas(config, url)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func commandMapb(config *pokeapi.LocationAreaResObject) error {
	var url string
	
	// If next is nil (aka first time using map command). We set it to the first page.
	if config.Previous == nil {
		fmt.Println("You're on the first page. Use the 'map' command to move forward!")
		return nil
	}

	// If not nil, we set the url to search for Next Page.
	if config.Previous != nil {	
		url = *config.Previous	
	}

	// Call API, pass in the URL to be searched
	err := pokeapi.GetLocationAreas(config, url)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func cleanInput(text string) []string {
	var result []string
	// Lowercase the text
	text = strings.ToLower(text)

	// Split on whitespace into a slice of strings. Spread them in order to append.
	result = append(result, strings.Fields(text)...)
	return result
}