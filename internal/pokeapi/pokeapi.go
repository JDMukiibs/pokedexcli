package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

// GetLocationAreas issues a GET to the specified endpoint that
// has the next or previous location areas using our Client.
// If the request is successful, continue to unmarshal
// and return a non-empty LocationAreas struct with a nil error.
// Otherwise, log any errors encountered along the flow
func (c *Client) GetLocationAreas(pageURL *string) (LocationAreas, error) {
	fullURL := baseURL + "/location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}

	// check our cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		fmt.Println("cache hit")
		locationAreasResponse := LocationAreas{}
		err := json.Unmarshal(dat, &locationAreasResponse)
		if err != nil {
			log.Println("Failed to unmarshal response from pokeapi. Try again later")
			return LocationAreas{}, err
		}

		return locationAreasResponse, nil
	}
	fmt.Println("cache miss")

	request, err := http.NewRequest("GET", fullURL, nil)
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

	// save to cache
	c.cache.Add(fullURL, body)

	return locationAreasResponse, nil
}

// GetLocationAreaDetails issues a GET to the specified endpoint that
// has detailed information about a specified location area using our Client.
// If the request is successful, continue to unmarshal
// and return a non-empty LocationAreaDetail struct with a nil error.
// Otherwise, log any errors encountered along the flow
func (c *Client) GetLocationAreaDetails(locationArea string) (LocationAreaDetail, error) {
	fullURL := baseURL + "/location-area/" + locationArea

	// check our cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		fmt.Println("cache hit")
		locationAreasResponse := LocationAreaDetail{}
		err := json.Unmarshal(dat, &locationAreasResponse)
		if err != nil {
			log.Println("Failed to unmarshal response from pokeapi. Try again later")
			return LocationAreaDetail{}, err
		}

		return locationAreasResponse, nil
	}
	fmt.Println("cache miss")

	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		log.Println(err)
		return LocationAreaDetail{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		log.Println(err)
		return LocationAreaDetail{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return LocationAreaDetail{}, err
	}

	locationAreaDetailResponse := LocationAreaDetail{}
	err = json.Unmarshal(body, &locationAreaDetailResponse)
	if err != nil {
		log.Println("Failed to unmarshal response from pokeapi. Try again later")
		return LocationAreaDetail{}, err
	}

	// save to cache, save fullUrl to struct
	c.cache.Add(fullURL, body)

	return locationAreaDetailResponse, nil
}
