package db

import (
	"fmt"

	"sync"

	es "github.com/DoeMoor/pokedexcli/internal/endpoint_scheme"
)

type pokedexDB struct {
	mu          sync.Mutex
	inventory   inventory
	pokemonList map[string]es.Pokemon
}

type inventory struct {
	pokemons map[string][]struct {
		name    string
		id      int
		deleted bool
	}
}

var db pokedexDB

// GetUserInventory returns the user's inventory
// it is a singleton
func GetUserInventory() *pokedexDB {
	if db.inventory.pokemons == nil {
		db.inventory.pokemons = make(map[string][]struct {
			name    string
			id      int
			deleted bool
		})
		db.pokemonList = make(map[string]es.Pokemon)
	}
	return &db
}

func (db *pokedexDB) AddPokemon(p es.Pokemon) {
	db.mu.Lock()
	defer db.mu.Unlock()
	newID := len(db.inventory.pokemons[p.Name])
	db.inventory.pokemons[p.Name] = append(db.inventory.pokemons[p.Name],
		struct {
			name    string
			id      int
			deleted bool
		}{
			name:    p.Name,
			id:      newID,
			deleted: false,
		})

	_, ok := db.pokemonList[p.Name]
	if !ok {
		db.pokemonList[p.Name] = p
	}
}

func (db *pokedexDB) RemovePokemon(name string, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	if len(db.inventory.pokemons[name]) == 0 {
		return fmt.Errorf("Pokemon %v not in inventory\n", name)
	}
	for index, pokemon := range db.inventory.pokemons[name] {
		if pokemon.id == id {
			db.inventory.pokemons[name][index].deleted = true
			fmt.Printf("Pokemon %v | %v removed from inventory\n", name, id)
			return nil
		}
	}
	return fmt.Errorf("Pokemon %v not in inventory\n", name)
}

func (db *pokedexDB) GetPokemonsListFromPokedex() (map[string]es.Pokemon ,error){
	db.mu.Lock()
	defer db.mu.Unlock()
	if db.pokemonList == nil {
		return nil, fmt.Errorf("No pokemons in inventory")
	}
	return db.pokemonList, nil
}

func (db *pokedexDB) GetPokemonFromPokedex(name string) (es.Pokemon, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	pokemon, ok := db.pokemonList[name]
	if !ok {
		return es.Pokemon{}, fmt.Errorf("Pokemon %v not in pokedex\n", name)
	}
	return pokemon, nil
}
