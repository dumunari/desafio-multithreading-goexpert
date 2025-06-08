package dto

type ViaCEP struct {
	Cep        string `json:"cep"`
	Localidade string `json:"localidade"`
	Estado     string `json:"estado"`
}

type BrasilApi struct {
	Cep        string `json:"cep"`
	Estado     string `json:"state"`
	Localidade string `json:"city"`
}
