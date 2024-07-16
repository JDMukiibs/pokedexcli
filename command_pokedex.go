package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, arguments []string) error {
	if len(cfg.pokedex) == 0 {
		return errors.New("you haven't caught any pokemon yet")
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
