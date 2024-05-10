package pokeapi

import (
	"io"
	"log"
	"net/http"
)

func getResource(endpoint string) ([]byte, error) {
	res, err := http.Get(endpoint)
	if err != nil {
		return []byte{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed\nstatus code: %d\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
