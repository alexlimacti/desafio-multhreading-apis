package api

type Address struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type APIResponse struct {
	Source  string
	Address *Address
	Error   error
}

type Fetcher interface {
	Fetch(cep string) (*Address, error)
}
