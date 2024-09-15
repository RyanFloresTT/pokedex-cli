package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/RyanFloresTT/pokedex-cli/internal"
)

func getLocations(url string) (location, error) {
	var location location

	cache := internal.NewCache(5)

	requestAPI(url, &location, cache)

	return location, nil
}

func getPokemonEncounters(location string) (AreaLocation, error) {
	var area AreaLocation

	cache := internal.NewCache(5)

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", 1)

	requestAPI(url, &area, cache)

	for _, res := range area.PokemonEncounters {
		fmt.Println(res.Pokemon.Name)
	}

	return area, nil
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

func getLocationID(name string) int {
	locations := getLocations()
}
