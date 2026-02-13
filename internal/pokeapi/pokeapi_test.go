package pokeapi

import (
	"testing"
	"time"

	"github.com/seanfinnessy/pokedexcli/internal/pokecache"
)

func TestGetPokemonInLocation(t *testing.T) {
	cache := pokecache.NewCache(5 * time.Second)
	tests := []struct {
		url            string
		expectedToPass bool
	}{
		{
			url:            LocationsURL + "ravaged-path-area",
			expectedToPass: true,
		},
		{
			url:            LocationsURL + "fake-area",
			expectedToPass: false,
		},
	}

	for _, test := range tests {
		testBool := GetPokemonInLocation(cache, test.url)
		if (testBool != nil) == test.expectedToPass {
			t.Errorf("Test failed for GetPokemonInLocation.")
		}
	}
}

func TestGetLocationAreas(t *testing.T) {
	cache := pokecache.NewCache(5 * time.Second)
	respObj := &LocationAreaResObject{}

	tests := []struct {
		url            string
		expectedToPass bool
	}{
		{
			url:            LocationsURL,
			expectedToPass: true,
		},
		{
			url:            "bad url",
			expectedToPass: false,
		},
	}

	for _, test := range tests {
		testBool := GetLocationAreas(respObj, cache, test.url)
		if (testBool != nil) == test.expectedToPass {
			t.Errorf("Test failed for GetPokemonInLocation.")
		}
	}
}
