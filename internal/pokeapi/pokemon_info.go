package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(pokemonName string) (Pokemon, error) {
	fullUrl := baseUrl + "/pokemon/" + pokemonName
	payload, ok := c.cache.Get(fullUrl)

	if ok {
		return getPokemonsByPayload(payload)
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullUrl, data)

	return getPokemonsByPayload(data)
}

func getPokemonsByPayload(payload []byte) (Pokemon, error) {
	responseLoc := Pokemon{}
	if err := json.Unmarshal(payload, &responseLoc); err != nil {
		return Pokemon{}, err
	}

	return responseLoc, nil
}