package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// TODO (Joshua): Update struct to keep track of starting id for a location-area
type pokeApiUrlTracker struct {
	previous string
	next     string
}

func main() {
	commands := getCommandMap()
	urlTracker := initializePokeApiUrlTracker()
	log.SetPrefix("Pokedex: ")
	log.SetFlags(0)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		userInput := getUserInput(scanner)
		command, ok := commands[userInput]
		for !ok {
			log.Println("Unrecognized command. Please enter a recognized command or enter 'help' to see available commands.")
			userInput = getUserInput(scanner)
			command, ok = commands[userInput]
		}
		err := command.callback(&urlTracker)
		if err != nil {
			log.Println(err)
		}
	}
}

func getUserInput(scanner *bufio.Scanner) string {
	for {
		fmt.Print("pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			if len(input) == 0 {
				continue
			}
			cleanedInput := strings.ToLower(strings.TrimSpace(input))
			words := strings.Fields(cleanedInput)
			return words[0]
		} else {
			inputError := scanner.Err()
			if inputError != nil {
				log.Println("Error reading input command. Please enter a recognized command or enter 'help' to see available commands.")
			}
		}
	}
}

func initializePokeApiUrlTracker() pokeApiUrlTracker {
	return pokeApiUrlTracker{
		previous: "",
		next:     "https://pokeapi.co/api/v2/location-area/",
	}
}
