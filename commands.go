package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type config struct {
	Next     *string
	Previous *string
}

var Pokedex map[string]Pokemon

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

func commandCatch(cfg *config, pName string) error {
	pokemon, err := GetPokemonInfo(pName)

	if err != nil {
		return err
	}

	fmt.Print("Throwing a ball at " + pokemon.Name)

	if tryCatchPokemon(pokemon.BaseExperience) {
		fmt.Println(pokemon.Name + " was caught!")
		Pokedex[pokemon.Name] = pokemon
		return nil
	}

	fmt.Println(pokemon.Name + " escaped!")
	return nil
}

func tryCatchPokemon(baseExp int) bool {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Print(".")
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println()
	randomNum := rand.Intn(101)

	catchChance := calculateCatchChance(baseExp)

	return randomNum <= int(catchChance)
}

func calculateCatchChance(baseExp int) float64 {
	baseFactor := 100.0
	experienceFactor := 20.0
	minCatchChance := 10.0

	catchChance := (baseFactor / (float64(baseExp)/experienceFactor + 1)) + minCatchChance

	if catchChance > 100 {
		catchChance = 100
	}

	return catchChance
}

func commandInspect(cfg *config, pokemonName string) error {
	if pokemon, exists := Pokedex[pokemonName]; exists {
		fmt.Println("Name: " + pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, element := range pokemon.Types {
			fmt.Printf(" - %s\n", element.Type.Name)
		}
		return nil
	}
	return errors.New("you have not caught that pokemon")
}

func commandPokedex(cfg *config, input string) error {
	if len(Pokedex) == 0 {
		return errors.New("pokedex is empty")
	}

	fmt.Println("Your Pokedex:")

	for _, pokemon := range Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
