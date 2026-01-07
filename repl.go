package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
		command.callback()
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println("")

	for _, value  := range getCommands() {
		helpMsg := fmt.Sprintf("%s: %s", value.name, value.description)
		fmt.Println(helpMsg)
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