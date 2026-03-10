package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
	"fmt"
	"github.com/rasqi7/pokedexcli/pokecache"
	"time"
)
var cache = pokecache.NewCache(5 * time.Minute)

type ApiResponse struct {
	Count int
	Next string
	Previous string
	Results []LocationArea
}

type LocationArea struct {
	Name string
	Url string
}

func GetLocationAreas(url string) (ApiResponse, error) {
	var data ApiResponse
	body, found := cache.Get(url)
	if found {
		err := json.Unmarshal(body, &data)
		if err != nil {
			return ApiResponse{}, err
		}

		return data, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return ApiResponse{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return ApiResponse{}, fmt.Errorf("non-OK HTTP status: %s", res.Status)
	}
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return ApiResponse{}, err
	}
	cache.Add(url, body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		return ApiResponse{}, err
	}

	return data, nil
}