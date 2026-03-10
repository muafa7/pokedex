package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
	"fmt"
)
const baseLocationURL = "https://pokeapi.co/api/v2/location-area/"

type LocationAreaResponse struct {
	PokemonEncounters []Encounter `json:"pokemon_encounters"`
}

type Encounter struct {
	Pokemon Pokemon
}

func GetLocationArea(areaName string) (LocationAreaResponse, error) {
	fullUrl := baseLocationURL + areaName
	var data LocationAreaResponse
	body, found := cache.Get(fullUrl)
	if found {
		err := json.Unmarshal(body, &data)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		return data, nil
	}
	res, err := http.Get(fullUrl)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return LocationAreaResponse{}, fmt.Errorf("non-OK HTTP status: %s", res.Status)
	}
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	cache.Add(fullUrl, body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return data, nil
}