package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	url := "https://pokeapi.co/api/v2/location-area/"

	config := config{
		Next:     &url,
		Previous: nil,
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		parts := strings.Split(input, " ")

		if len(parts) > 2 {
			fmt.Println("Error: Invalid command format. Use '<command> [parameter]'.")
			continue
		}

		command := parts[0]
		var parameter string
		if len(parts) == 2 {
			parameter = parts[1]
		}

		if cmd, exists := commands[command]; exists {
			err := cmd.callback(&config, parameter)
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
		} else {
			fmt.Println(command + " " + parameter)
			fmt.Println("Unknown command. Type 'help' for available commands.")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}
