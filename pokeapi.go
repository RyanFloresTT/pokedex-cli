package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/RyanFloresTT/pokedex-cli/internal"
)

func getLocations(url string) (location, error) {
	var location location

	cache := internal.NewCache(5)

	err := requestAPI(url, &location, cache)
	if err != nil {
		return location, err
	}

	return location, nil
}

func getPokemonEncounters(location string) (AreaLocation, error) {
	var area AreaLocation

	cache := internal.NewCache(5)

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)

	err := requestAPI(url, &area, cache)
	if err != nil {
		return area, err
	}

	return area, nil
}

func GetPokemonInfo(name string) (Pokemon, error) {
	var pokemon Pokemon

	if name == "" {
		return pokemon, errors.New("must specify a pokemon to catch")
	}

	cache := internal.NewCache(5)

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)

	err := requestAPI(url, &pokemon, cache)
	if err != nil {
		return pokemon, err
	}

	if pokemon.ID == 0 {
		return pokemon, errors.New("pokemon name not found")
	}

	return pokemon, nil
}

func requestAPI(url string, dataStruct any, cache *internal.Cache) error {
	if data, exists := cache.Get(url); exists {
		json.Unmarshal(data, &dataStruct)
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	cache.Add(url, data)

	json.Unmarshal(data, &dataStruct)

	return nil
}
