package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jdmukiibs/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeApiClient pokeapi.Client
	previousUrl   *string
	nextUrl       *string
}

func main() {
	commands := getCommandMap()
	pokeApiClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeApiClient: pokeApiClient,
	}
	log.SetPrefix("Pokedex: ")
	log.SetFlags(0)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		userInput := getUserInput(scanner)
		command, ok := commands[userInput[0]]
		for !ok {
			log.Println("Unrecognized command. Please enter a recognized command or enter 'help' to see available commands.")
			userInput = getUserInput(scanner)
			command, ok = commands[userInput[0]]
		}
		err := command.callback(cfg, userInput[1:])
		if err != nil {
			log.Println(err)
		}
	}
}

// This method returns a list of strings that represent all the user has entered
// e.g. explore pastoria-city-area -> [explore, pastoria-city-area]
func getUserInput(scanner *bufio.Scanner) []string {
	for {
		fmt.Print("pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			if len(input) == 0 {
				continue
			}
			cleanedInput := strings.ToLower(strings.TrimSpace(input))
			words := strings.Fields(cleanedInput)
			return words
		} else {
			inputError := scanner.Err()
			if inputError != nil {
				log.Println("Error reading input command. Please enter a recognized command or enter 'help' to see available commands.")
			}
		}
	}
}
