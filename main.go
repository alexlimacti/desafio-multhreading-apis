package main

import (
	"fmt"
	"time"

	"github.com/alexl/desafio-multhreading-apis/internal/api"
)

func fetch(fetcher api.Fetcher, currentCEP string, sourceName string, ch chan<- api.APIResponse) {
	address, err := fetcher.Fetch(currentCEP)
	ch <- api.APIResponse{Source: sourceName, Address: address, Error: err}
}

func main() {
	cep := "48609070"
	ch := make(chan api.APIResponse)

	brasilAPI := &api.BrasilAPI{}
	viaCEP := &api.ViaCEP{}

	go fetch(brasilAPI, cep, "BrasilAPI", ch)
	go fetch(viaCEP, cep, "ViaCEP", ch)

	timeout := time.After(1 * time.Second)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch:
			if msg.Error == nil {
				fmt.Printf("Start: %s\n", cep)
				fmt.Printf("Received response from: %s\n", msg.Source)
				// Using %+v to print fields of the Address struct
				fmt.Printf("Address: %+v\n", msg.Address)
				return // Found a successful response
			} else {
				// Log the error but continue waiting for the other API
				fmt.Printf("Error from %s: %v\n", msg.Source, msg.Error)
			}
		case <-timeout:
			fmt.Println("Timeout: Request took longer than 1 second")
			return
		}
	}
}
