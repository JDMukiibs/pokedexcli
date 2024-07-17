package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/chzyer/readline"
	"github.com/jdmukiibs/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeApiClient pokeapi.Client
	previousUrl   *string
	nextUrl       *string
	pokedex       map[string]pokeapi.PokemonData
}

func main() {
	// Create a readline instance with customized configuration
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "pokedex > ",
		HistoryFile:     "/tmp/readline.tmp", // Enables persistence of commands
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		log.Fatalf("could not create readline instance: %s", err)
		os.Exit(0)
	}
	defer rl.Close()
	commands := getCommandMap()
	pokeApiClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeApiClient: pokeApiClient,
		pokedex:       make(map[string]pokeapi.PokemonData),
	}
	log.SetPrefix("Pokedex: ")
	log.SetFlags(0)

	for {
		userInput := getUserInputWithReadline(rl)
		command, ok := commands[userInput[0]]
		for !ok {
			log.Println("Unrecognized command. Please enter a recognized command or enter 'help' to see available commands.")
			userInput := getUserInputWithReadline(rl)
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
func getUserInputWithReadline(rl *readline.Instance) []string {
	for {
		line, err := rl.Readline()
		if err != nil { // Exit on interrupt signal (Ctrl+C) or EOF signal
			return []string{"exit"}
		}
		if len(line) == 0 {
			log.Println("Error reading input command. Please enter a recognized command or enter 'help' to see available commands.")
			continue
		}
		cleanedInput := strings.ToLower(strings.TrimSpace(line))
		words := strings.Fields(cleanedInput)
		return words
	}
}
