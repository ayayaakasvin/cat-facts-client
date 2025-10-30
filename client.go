package client

import (
	"net/http"
	"time"
)

const baseURL string = "https://meowfacts.herokuapp.com/"
const defaultTimeout = time.Second * 10

type Client struct {
	httpclient *http.Client
}

func NewDefaultClient() *Client {
	return &Client{
		httpclient: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

func NewClient(httpclient *http.Client) *Client {
	return &Client{httpclient}
}
