package pokecache

import (
	"testing"
	"time"
)

func TestAddToCache(t *testing.T) {
	cache := NewCache(5 * time.Second)

	tests := []struct {
		url string
		value []byte
		expected int
	} {
		{
			url: "test url",
			value: make([]byte, 5),
			expected: 1,
		},
		{
			url: "test url2",
			value: make([]byte, 6),
			expected: 2,
		},
	}

	for _, testCase := range(tests) {
		cache.Add(testCase.url, testCase.value)

		if len(cache.cacheEntry) != testCase.expected {
			t.Errorf("Expected cache to have length of %d, instead got: %d", testCase.expected, len(cache.cacheEntry))
		}
	}
}

func TestGetCacheValue(t *testing.T) {
	cache := NewCache(3 * time.Second)
	byteSlice1 := make([]byte, 1)
	byteSlice2 := make([]byte, 3)

	tests := []struct {
		url string
		value []byte
		foundValue bool
		
	} {
		{
			url: "testurlkey",
			value: byteSlice1,
			foundValue: true,
			
		},
		{
			url: "testurlkey2",
			value: byteSlice2,
			foundValue: true,
			
		},
	}

	for _, testCase := range(tests) {
		cache.Add(testCase.url, testCase.value)
		returnVal, ok := cache.Get(testCase.url);
		if !ok {
			t.Errorf("Unable to locate key %s in cache", testCase.url)
		}

		if len(returnVal) != len(testCase.value) {
			t.Errorf("Returned wrong value from cache")
		}
	}
}