package pokeapi

import (
	"io"
	"log"
	"net/http"
)

func getResource(endpoint string) []byte {
	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed\nstatus code: %d\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	return body
}
