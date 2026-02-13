package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaResObject struct {
	Previous *string    `json:"previous"`
	Next     *string    `json:"next"`
	Results  []Location `json:"results"`
}

type PokemonResObject struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Cache interface {
	Get(key string) ([]byte, bool)
	Add(key string, val []byte)
}

func GetPokemonInLocation(cache Cache, url string) error {
	if len(url) == 0 {
		return fmt.Errorf("Empty url")
	}

	bytesToUnmarshal, err := GetBytes(cache, url)
	if err != nil {
		return fmt.Errorf("Issue with GetBytes")
	}

	// unmarshal and display
	var pokemonResObject *PokemonResObject
	errUnmarshal := json.Unmarshal(bytesToUnmarshal, &pokemonResObject)
	if errUnmarshal != nil {
		return fmt.Errorf("Issue unmarshalling pokemon json: %w", errUnmarshal)
	}

	ListPokemon(pokemonResObject)
	return nil
}

func GetLocationAreas(respObj *LocationAreaResObject, cache Cache, url string) error {
	if len(url) == 0 {
		return fmt.Errorf("Empty url")
	}

	bytesToUnmarshal, err := GetBytes(cache, url)
	if err != nil {
		return fmt.Errorf("Issue with GetBytes")
	}

	// unmarshal and display
	errUnmarshal := json.Unmarshal(bytesToUnmarshal, respObj)
	if errUnmarshal != nil {
		return fmt.Errorf("Issue unmarshalling locations json: %w", errUnmarshal)
	}
	ListLocations(respObj)
	return nil
}

func ListLocations(responseBody *LocationAreaResObject) {
	// extract locations
	var results []string
	locations := responseBody.Results
	for _, locationObj := range locations {
		results = append(results, locationObj.Name)
	}

	for _, location := range locations {
		fmt.Println(location.Name)
	}
	fmt.Println()

}

func ListPokemon(p *PokemonResObject) {
	// List pokemon
	for _, encounter := range p.PokemonEncounters {
		fmt.Println("- " + encounter.Pokemon.Name)
	}
	fmt.Println()
}

func GetBytes(cache Cache, url string ) ([]byte, error) {
	var bytesToUnmarshal []byte
	if cachedBytes, isCached := cache.Get(url); isCached {
		bytesToUnmarshal = cachedBytes
	} else {
		res, errGet := http.Get(url)
		if errGet != nil {
			return bytesToUnmarshal, fmt.Errorf("Issue calling API: %s", url)
		}

		// decode json response into config which our main REPL loop uses to navigate
		if bodyBytes, errReadAll := io.ReadAll(res.Body); errReadAll == nil {
			bytesToUnmarshal = bodyBytes
		}

		// add bytes to cache
		cache.Add(url, bytesToUnmarshal)
	}

	return bytesToUnmarshal, nil
}
