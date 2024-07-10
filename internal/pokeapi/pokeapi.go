package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

// LocationAreas represents the HTTP response from the PokeAPI
// when requesting for location areas of pokemon
type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// GetLocationAreas issues a GET to the specified endpoint that
// has the next or previous location areas. If the response is
// status code is > 299, an empty LocationAreas struct is
// returned alongside an error. Otherwise, continue to unmarshal
// and return a non-empty LocationAreas struct with a nil error.
func GetLocationAreas(url string) (LocationAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return LocationAreas{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return LocationAreas{}, errors.New("response status code out of expected range. try again later")
	}
	if err != nil {
		log.Println(err)
		return LocationAreas{}, err
	}
	apiResponse := LocationAreas{}
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Println("Failed to unmarshal response from pokeapi. Try again later")
		return LocationAreas{}, err
	}

	return apiResponse, nil
}
