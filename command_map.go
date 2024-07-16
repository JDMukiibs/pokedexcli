package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, arguments []string) error {
	locationAreas, err := cfg.pokeApiClient.GetLocationAreas(cfg.nextUrl)
	if err != nil {
		return err
	}
	// Update urlTracker to have a new next and previous
	cfg.nextUrl = locationAreas.Next
	cfg.previousUrl = locationAreas.Previous
	// Print out our location areas
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapBack(cfg *config, arguments []string) error {
	if cfg.previousUrl == nil {
		return errors.New("currently on first page of results. no previous results to show")
	}

	locationAreas, err := cfg.pokeApiClient.GetLocationAreas(cfg.previousUrl)
	if err != nil {
		return err
	}
	// Update urlTracker to have a new next and previous
	cfg.nextUrl = locationAreas.Next
	cfg.previousUrl = locationAreas.Previous
	// Print out our location areas
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	return nil
}
