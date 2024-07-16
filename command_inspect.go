package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, arguments []string) error {
	if len(arguments) != 1 {
		return errors.New("only one pokemon name should be passed in with command")
	}
	pokemon, ok := cfg.pokedex[arguments[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	// Print details of the Pokemon
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	stats := pokemon.Stats
	fmt.Println("Stats:")
	for _, s := range stats {
		fmt.Printf("  - %s: %d\n", s.Stat.Name, s.BaseStat)
	}
	types := pokemon.Types
	fmt.Println("Types:")
	for _, t := range types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}
