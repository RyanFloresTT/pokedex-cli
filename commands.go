package main

import (
	"fmt"
	"os"
)

var pagination int

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandHelp() error {
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

func commandMap() error {
	return nil
}

func commandMapb() error {
	return nil
}
