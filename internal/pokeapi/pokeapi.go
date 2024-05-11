package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/wtwingate/pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2/"

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient() *Client {
	client := Client{
		httpClient: http.Client{
			Timeout: 30 * time.Second,
		},
		cache: *pokecache.NewCache(5 * time.Minute),
	}

	return &client
}

func (c *Client) getResource(pageURL string) ([]byte, error) {

	req, err := http.NewRequest("GET", pageURL, nil)
	if err != nil {
		return []byte{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	if res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	return body, err
}
