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

type Cache interface{
	Get(key string) ([]byte, bool)
	Add(key string, val []byte)
}

func GetLocationAreas(respObj *LocationAreaResObject, cache Cache, url string) error {
	if len(url) == 0 {
		return fmt.Errorf("Empty url")
	}

	var bytesToUnmarshal []byte
	if cachedBytes, isCached := cache.Get(url); isCached {
		bytesToUnmarshal = cachedBytes
	} else {
		res, errGet := http.Get(url)
		if errGet != nil {
			return fmt.Errorf("Issue retrieving locations: %w", errGet)
		}

		// decode json response into config which our main REPL loop uses to navigate
		bodyBytes, errReadAll := io.ReadAll(res.Body)
		if errReadAll != nil {
			return fmt.Errorf("Issue reading bytes: %w", errReadAll)
		}
		bytesToUnmarshal = bodyBytes

		// add bytes to cache
		cache.Add(url, bodyBytes)
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

}
