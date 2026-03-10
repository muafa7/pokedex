package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
	"fmt"
)

const basePokemonURL = "https://pokeapi.co/api/v2/pokemon"

type Pokemon struct {
	Name           string
	BaseExperience int      `json:"base_experience"`
	Height         int      `json:"height"`
	Weight         int      `json:"weight"`
	Stats          []Stat   `json:"stats"`
	Types          []Type   `json:"types"`
}

type Stat struct {
	BaseStat int      `json:"base_stat"`
	Stat     StatInfo `json:"stat"`
}

type StatInfo struct {
	Name string `json:"name"`
}

type Type struct {
	Type TypeInfo `json:"type"`
}

type TypeInfo struct {
	Name string `json:"name"`
}

func GetPokemon(name string) (Pokemon, error) {
	fullUrl := basePokemonURL + "/" + name

	var data Pokemon
	body, found := cache.Get(fullUrl)
	if found {
		err := json.Unmarshal(body, &data)
		if err != nil {
			return Pokemon{}, err
		}

		return data, nil
	}
	res, err := http.Get(fullUrl)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("non-OK HTTP status: %s", res.Status)
	}
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	cache.Add(fullUrl, body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		return Pokemon{}, err
	}

	return data, nil
}