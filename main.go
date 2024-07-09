package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var exitPokedex bool
	log.SetPrefix("Pokedex: ")
	log.SetFlags(0)
	scanner := bufio.NewScanner(os.Stdin)

	for !exitPokedex {
		fmt.Print("pokedex > ")
		scanner.Scan()
		inputError := scanner.Err()
		for inputError != nil {
			log.Println("Error reading input command. Please enter a recognized command or enter 'help' to see available commands.")
			fmt.Print("pokedex > ")
			scanner.Scan()
			inputError = scanner.Err()
		}
		command := strings.TrimSpace(scanner.Text())
		if command == "exit" {
			exitPokedex = true
			continue
		}
		fmt.Println("You entered: ", command)
	}
}
