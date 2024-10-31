package pokeapi

import (
	"net/http"
	"time"

	"github.com/st5/pokedexcli/internal/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

