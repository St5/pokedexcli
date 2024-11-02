package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreaList, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint

	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	//check a cache
	payload, ok := c.cache.Get(fullUrl)
	if ok {
		responseLoc := LocationAreaList{}

		if err := json.Unmarshal(payload, &responseLoc); err != nil {
			return LocationAreaList{}, err
		}

		return responseLoc, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreaList{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreaList{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreaList{}, fmt.Errorf("bad status %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaList{}, err
	}

	responseLoc := LocationAreaList{}

	if err = json.Unmarshal(data, &responseLoc); err != nil {
		return LocationAreaList{}, err
	}

	c.cache.Add(fullUrl, data)

	return responseLoc, nil

}
