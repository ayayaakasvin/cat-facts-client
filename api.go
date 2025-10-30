package client

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func (c *Client) FunFact() (string, error) {
	req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.httpclient.Do(req)
	if err != nil {
		return "", err
	} else if resp.StatusCode >= 400 && resp.StatusCode <= 500 {
		return "", fmt.Errorf("error status code: %d %s", resp.StatusCode, resp.Status)
	}

	defer resp.Body.Close()

	var cfp CatFactResponse
	if err := json.NewDecoder(resp.Body).Decode(&cfp); err != nil {
		return "", err
	}

	return cfp.Data[0], nil
}