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
	// TODO: getting an issue 
	var url string
	if config.Next != nil {		
		fmt.Println(*config.Next)
	}
	fmt.Println(config.Previous)

	if config.Next == nil {
		url = "https://pokeapi.co/api/v2/location-area/"
		// config.Next is a pointer to string, so give it the url address
		config.Next = &url
	}

	fmt.Println(url)
	err := pokeapi.GetLocationAreas(url)
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