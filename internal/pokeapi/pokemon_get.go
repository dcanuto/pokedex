package pokeapi

import (
	"encoding/json"

	"github.com/dcanuto/pokedexcli/internal/pokecache"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseUrl + "pokemon/" + pokemonName

	body, err := pokecache.GetFromOrAddToCache(url, &c.Cache)
	if err != nil {
		return Pokemon{}, err
	}

	result := Pokemon{}
	err = json.Unmarshal(body, &result)
	return result, err
}
