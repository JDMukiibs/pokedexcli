package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jdmukiibs/internal/pokeapi"
)

type config struct {
	pokeApiClient pokeapi.Client
	previousUrl   *string
	nextUrl       *string
}

func main() {
	commands := getCommandMap()
	pokeApiClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeApiClient: pokeApiClient,
	}
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
		err := command.callback(cfg)
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
