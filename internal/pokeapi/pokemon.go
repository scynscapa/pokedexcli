package pokeapi

import (
	"fmt"
	"io"
	"encoding/json"
	"net/http"
	"math/rand"
)

type Pokemon struct {
	Name		string		`json:"name"`
	BaseExp		int			`json:"base_experience"`
	Height		int			`json:"height"`
	Weight		int			`json:"weight"`
	Stats		[]PokeStats	`json:"stats"`
	Types		[]Types		`json:"types"`
}

type PokeStats struct {
	BaseStat	int			`json:"base_stat"`
	Stat		Stat		`json:"stat"`
}

type Stat struct {
	Name		string		`json:"name"`
}

type Types struct {
	Type		Type		`json:"type"`
}

type Type struct {
	Name		string		`json:"name"`
}

func CommandInspect(config *ConfigStruct, toInspect string) error {
	pokemon, exists := config.Pokedex[toInspect]
	if !exists {
		fmt.Println("You have not caught that Pokemon")
		return nil
	}

	height := pokemon.Height
	weight := pokemon.Weight

	fmt.Println("Name: " + pokemon.Name)
	fmt.Printf("Height: %d\nWeight: %d\nStats:\n", height, weight)

	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, types := range pokemon.Types {
		fmt.Printf("  - %s\n", types.Type.Name)
	}

	return nil
}

func NewPokedex() map[string]Pokemon {
	pokedex := make(map[string]Pokemon)

	return pokedex
}

func CommandCatch(config *ConfigStruct, toCapture string) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + toCapture

	pokemon, err := getPokemon(config, url)
	if err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at " + pokemon.Name + "...")

	expModMax := int(((-float64((pokemon.BaseExp - 609)))*0.15) + 11)
	expModMin := int(((-float64((pokemon.BaseExp - 609)))*0.09))

	chance := rand.Intn(100)

	if chance < expModMax && chance > expModMin {
		config.Pokedex[pokemon.Name] = pokemon
		fmt.Println("Caught " + pokemon.Name + "!")
	}
	return nil
}

func getPokemon(config *ConfigStruct, url string) (Pokemon, error) {
	cached, exists := config.Cache.Get(url)
	body := []byte{}
	if exists {
		body = cached
	} else {
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
	}

	var data Pokemon

	err := json.Unmarshal(body, &data)
	if err != nil {
		return Pokemon{}, err
	}

	return data, nil
}