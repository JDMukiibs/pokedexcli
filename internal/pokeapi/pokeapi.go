package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
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
// has the next or previous location areas using our Client. 
// If the request is successful, continue to unmarshal
// and return a non-empty LocationAreas struct with a nil error.
// Otherwise, log any errors encountered along the flow
func (c *Client) GetLocationAreas(pageURL *string) (LocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return LocationAreas{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		log.Println(err)
		return LocationAreas{}, err
	}
	defer response.Body.Close()

	
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return LocationAreas{}, err
	}

	locationAreasResponse := LocationAreas{}
	err = json.Unmarshal(body, &locationAreasResponse)
	if err != nil {
		log.Println("Failed to unmarshal response from pokeapi. Try again later")
		return LocationAreas{}, err
	}

	return locationAreasResponse, nil
}
