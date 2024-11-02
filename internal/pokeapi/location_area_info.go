package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationInfo(pathUrl string) (LocationArea, error) {

	fullUrl := baseUrl + "/location-area/" + pathUrl

	payload, ok := c.cache.Get(fullUrl)

	if ok {
		return getLocationOfPokemons(payload)
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullUrl, data)
	return getLocationOfPokemons(data)
}

func getLocationOfPokemons(payload []byte) (LocationArea, error) {
	responseLoc := LocationArea{}
	if err := json.Unmarshal(payload, &responseLoc); err != nil {
		return LocationArea{}, err
	}

	return responseLoc, nil
}
