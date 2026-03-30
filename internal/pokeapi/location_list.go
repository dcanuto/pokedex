package pokeapi

import (
	"encoding/json"

	"github.com/dcanuto/pokedexcli/internal/pokecache"
)

func (c *Client) GetLocationAreas(pageUrl *string) (locationResourceList, error) {
	url := baseUrl + "location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	body, err := pokecache.GetFromOrAddToCache(url, &c.Cache)
	if err != nil {
		return locationResourceList{}, err
	}

	m := locationResourceList{}
	err = json.Unmarshal(body, &m)

	return m, err
}
