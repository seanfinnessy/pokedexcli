package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
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

type PokemonStat struct {
	BaseStatValue int `json:"base_stat"`
	Stat Name `json:"stat"`
}

type PokemonType struct {
	Type Name `json:"type"`
}

type Name struct {
	Name string `json:"name"`
}


type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	BaseExperience int `json:"base_experience"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []PokemonStat `json:"stats"`
	Types []PokemonType `json:"types"`
	
}

type Cache interface {
	Get(key string) ([]byte, bool)
	Add(key string, val []byte)
}



func GetPokemonStats(cache Cache, pokedex map[string]Pokemon,  url string) error {
	if len(url) == 0 {
		return fmt.Errorf("Empty url")
	}

	bytesToUnmarshal, err := GetBytes(cache, url)
	if err != nil {
		return err
	}

	var pokemonResponse *Pokemon
	if err := json.Unmarshal(bytesToUnmarshal, &pokemonResponse); err != nil {
		return fmt.Errorf("Error unmarshalling pokemon bytes: %w." ,err )
	}


	CatchPokemon(pokemonResponse, pokedex)
	return nil

}

func CatchPokemon(pokemonResponse *Pokemon, pokedex map[string]Pokemon) {
	name := pokemonResponse.Name
	fmt.Println("Throwing a Pokeball at " + name + "...")
	baseExperience := float64(pokemonResponse.BaseExperience)
	probabilityOfCatch := math.Pow((1.0 - (baseExperience / 608.0)), 3)
	if rand.Float64() < probabilityOfCatch {
		fmt.Println(name + " was caught!")
		fmt.Println("You may now inspect it with the inspect command.")
		pokedex[name] = *pokemonResponse
	} else {
		fmt.Println(name + " escaped!")
	}
}

func GetPokemonInLocation(cache Cache, url string) error {
	if len(url) == 0 {
		return fmt.Errorf("Empty url")
	}

	bytesToUnmarshal, err := GetBytes(cache, url)
	if err != nil {
		return err
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
		return err
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
		if errGet != nil || res.StatusCode != 200 {
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
