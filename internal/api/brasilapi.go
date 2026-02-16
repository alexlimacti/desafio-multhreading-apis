package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type BrasilAPI struct{}

type BrasilAPIResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func (b *BrasilAPI) Fetch(cep string) (*Address, error) {
	start := time.Now()
	url := "https://brasilapi.com.br/api/cep/v1/" + cep
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data BrasilAPIResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	fmt.Printf("BrasilAPI duration: %v\n", time.Since(start))

	return &Address{
		Cep:          data.Cep,
		State:        data.State,
		City:         data.City,
		Neighborhood: data.Neighborhood,
		Street:       data.Street,
		Service:      data.Service,
	}, nil
}
