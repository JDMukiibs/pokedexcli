package pokeapi

import "fmt"

// TODO (Joshua): Implement intended GET request with pokeapi-go pkg
func GetLocations(url string) (string, error) {
	return fmt.Sprintf("URL: %s", url), nil	
}
