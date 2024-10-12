package db

import (
	"fmt"

	es "github.com/DoeMoor/pokedexcli/internal/endpoint_scheme"
	"sync"
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

func (db *pokedexDB) RemovePokemon(name string, id int) {
	db.mu.Lock()
	defer db.mu.Unlock()
	if len(db.inventory.pokemons[name]) == 0 {
		fmt.Printf("Pokemon %v not in inventory\n", name)
		return
	}
	for index, pokemon := range db.inventory.pokemons[name] {
		if pokemon.id == id {
			db.inventory.pokemons[name][index].deleted = true
			fmt.Printf("Pokemon %v | %v removed from inventory\n", name, id)
			return
		}
	}
}

func (db *pokedexDB) GetPokemonList() {
	db.mu.Lock()
	defer db.mu.Unlock()
	fmt.Println("Pokemons in inventory:")
	for _, pokemons := range db.inventory.pokemons {
		for _, pokemon := range pokemons {
			if !pokemon.deleted {
				fmt.Printf(" - %v | %v\n", pokemon.name, pokemon.id)
			}
		}
	}
}

func (db *pokedexDB) GetPokemonFromInventory(name string, id int) {
	db.mu.Lock()
	defer db.mu.Unlock()
	if len(db.inventory.pokemons[name]) == 0 {
		fmt.Printf("Pokemon %v not in inventory\n", name)
		return
	}
	pokemons := db.inventory.pokemons[name]
	for _, pokemon := range pokemons {
		if pokemon.id == id {
			fmt.Printf("Pokemon %v found in inventory\n %v", name, db.pokemonList[name].Abilities)
			return
		}
	}
	fmt.Printf("Pokemon %v not in inventory\n", name)

}
