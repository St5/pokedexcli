package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreaRes, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint

	if(pageUrl != nil){
		fullUrl = *pageUrl
	}

	//check a cache
	payload, ok := c.cache.Get(fullUrl)
	if ok {
		responseLoc := LocationAreaRes {}

		if err := json.Unmarshal(payload, &responseLoc); err != nil {
			return LocationAreaRes{}, err
		}

		return responseLoc, nil
	}

	req,err := http.NewRequest("GET", fullUrl, nil)
	if(err != nil){
		return LocationAreaRes{}, err
	}

	res, err := c.httpClient.Do(req)

	if(err != nil){
		return LocationAreaRes{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreaRes{}, fmt.Errorf("bad status %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if(err != nil){
		return LocationAreaRes{}, err
	}

	responseLoc := LocationAreaRes {}

	if err = json.Unmarshal(data, &responseLoc); err != nil {
		return LocationAreaRes{}, err
	}

	c.cache.Add(fullUrl, data)

	return responseLoc, nil

}