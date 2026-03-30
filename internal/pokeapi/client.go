package pokeapi

import (
	"time"

	"github.com/dcanuto/pokedexcli/internal/pokecache"
)

type Client struct {
	Cache pokecache.Cache
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		Cache: pokecache.NewCache(cacheInterval),
	}
}
