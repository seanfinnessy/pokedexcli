package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	pokeapi "github.com/seanfinnessy/pokedexcli/internal/pokeapi"
	pokecache "github.com/seanfinnessy/pokedexcli/internal/pokecache"
)

// Create a struct holding pointer to our resp object and cache
type AppState struct {
	LocationResponse *pokeapi.LocationAreaResObject
	Cache *pokecache.Cache
}

func startRepl() {
	// Create a new scanner
	scanner := bufio.NewScanner(os.Stdin)

	// Initialize state as a pointer. We dont want to keep copying it
	state := &AppState {
		LocationResponse: &pokeapi.LocationAreaResObject{},
    	Cache: pokecache.NewCache(5 * time.Second),
	}


	// REPL
	for {
		fmt.Print("Pokedex > ")
		// Scan for input
		scanner.Scan()
		// Apply input to variable
		input := scanner.Text()
		// Clean input, grab first word, set command and show user
		cleanedInput := cleanInput(input)

		// Make sure command exists
		if len(cleanedInput) > 0 {
			command := cleanedInput[0]
			// Verify command
			checkCommand(command, state)
		}
	}
}

func checkCommand(commandString string, appState *AppState) {
	command, ok := getCommands()[commandString]
	if !ok {
		fmt.Println("Unknown command.")
	} else {
		// callback function, pass addr to config
		command.callback(appState)
	}
}

func commandExit(appState *AppState) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(appState *AppState) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println("")

	for _, value  := range getCommands() {
		helpMsg := fmt.Sprintf("%s: %s", value.name, value.description)
		fmt.Println(helpMsg)
	}
	return nil
}

func commandMap(appState *AppState) error {
	var url string
	
	// If next is nil (aka first time using map command). We set it to the first page.
	if appState.LocationResponse.Next == nil {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	// If not nil, we set the url to search for Next Page.
	if appState.LocationResponse.Next != nil {	
		url = *appState.LocationResponse.Next	
	}

	// Call API, pass in the URL to be searched and ptr the response object
	err := pokeapi.GetLocationAreas(appState.LocationResponse, appState.Cache, url)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func commandMapb(appState *AppState) error {
	var url string
	
	// If next is nil (aka first time using map command). We set it to the first page.
	if appState.LocationResponse.Previous == nil {
		fmt.Println("You're on the first page. Use the 'map' command to move forward!")
		return nil
	}

	// If not nil, we set the url to search for Next Page.
	if appState.LocationResponse.Previous != nil {	
		url = *appState.LocationResponse.Previous
	}

	// Call API, pass in the URL to be searched
	err := pokeapi.GetLocationAreas(appState.LocationResponse, appState.Cache, url)
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