package client

import (
	"fmt"
	"net/http"
	"encoding/json"
)

const baseURL string = "https://meowfacts.herokuapp.com/"

func FunFact() (string, error) {
	resp, err := http.Get(baseURL)
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