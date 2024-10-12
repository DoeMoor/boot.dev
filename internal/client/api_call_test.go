package client

import (
	"reflect"
	"testing"

	scheme "github.com/DoeMoor/pokedexcli/internal/endpoint_scheme"
)

type pokemon struct {
	Name string `json:"name"`
}

func TestApiCall(t *testing.T) {
	var loc scheme.Locations
	var resultLoc scheme.Locations
	var poke pokemon
	var resultPoke pokemon
	var pokemonLocationApiUrl string = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=5"

	t.Run(t.Name()+" matched scheme", func(t *testing.T) {
		err := ApiCall(pokemonLocationApiUrl, &loc)
		if err != nil {
			t.Errorf("ApiCall() error = %v", err)
			return
		}
		if reflect.TypeOf(loc) != reflect.TypeOf(resultLoc) {
			t.Errorf("ApiCall() = %v, want %v", loc, resultLoc)
			return
		}
	})

	t.Run(t.Name()+" unmatched scheme", func(t *testing.T) {
		err := ApiCall(pokemonLocationApiUrl, &poke)
		if err != nil {
			t.Errorf("ApiCall() error = %v", err)
			return
		}
		if reflect.TypeOf(poke) == reflect.TypeOf(resultLoc) {
			t.Errorf("ApiCall() = %v, want %v", poke, resultPoke)
			return
		}
	})
}
