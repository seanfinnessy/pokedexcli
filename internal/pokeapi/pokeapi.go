package pokeapi

import (
	"encoding/json"
	"fmt"
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

func GetLocationAreas(config *LocationAreaResObject, url string) error {
	if len(url) == 0 {
		return fmt.Errorf("Empty url")
	}

	res, errGet := http.Get(url)
	if errGet != nil {
		return fmt.Errorf("Issue retrieving locations: %w", errGet)
	}

	// decode json response into config which our main REPL loop uses to navigate
	decoder := json.NewDecoder(res.Body)
	err := decoder.Decode(config)
	if err != nil {
		return fmt.Errorf("Issue decoding locations json: %w", err)
	}

	// List all locations
	ListLocations(config)
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
		fmt.Println("------------------------------------")
	}

}
