package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ViaCEP struct{}

type ViaCEPResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (v *ViaCEP) Fetch(cep string) (*Address, error) {
	start := time.Now()
	url := "http://viacep.com.br/ws/" + cep + "/json/"
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

	var data ViaCEPResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	fmt.Printf("ViaCEP duration: %v\n", time.Since(start))

	return &Address{
		Cep:          data.Cep,
		State:        data.Uf,
		City:         data.Localidade,
		Neighborhood: data.Bairro,
		Street:       data.Logradouro,
		Service:      "via-cep",
	}, nil
}
