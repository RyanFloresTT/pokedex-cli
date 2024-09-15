package main

import (
	"errors"
	"fmt"
	"os"
)

type config struct {
	Next     *string
	Previous *string
}

func commandExit(cfg *config, input string) error {
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, input string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()

	return nil
}

func commandMap(cfg *config, input string) error {
	if cfg.Next == nil {
		fmt.Println("Can't page any further!")
		return nil
	}

	loc, err := getLocations(*cfg.Next)
	if err != nil {
		return err
	}

	for _, res := range loc.Results {
		fmt.Println(res.Name)
	}

	cfg.Next = loc.Next
	cfg.Previous = loc.Previous

	return nil
}

func commandMapb(cfg *config, input string) error {
	if cfg.Previous == nil {
		fmt.Println("Can't page back any further!")
		return nil
	}

	loc, err := getLocations(*cfg.Previous)
	if err != nil {
		return err
	}

	for _, res := range loc.Results {
		fmt.Println(res.Name)
	}

	cfg.Next = loc.Next
	cfg.Previous = loc.Previous

	return nil
}

func commandExplore(cfg *config, location string) error {
	if location == "" {
		return errors.New("need to specify an area to explore")
	}

	loc, err := getPokemonEncounters(location)
	if err != nil {
		return err
	}

	fmt.Println("Exploring " + location + "...")

	fmt.Println("Found Pokemon:")

	for _, res := range loc.PokemonEncounters {
		fmt.Printf(" - %s\n", res.Pokemon.Name)
	}

	return nil
}
