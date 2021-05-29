package model

type Country struct {
	CountryName string   `json:"name"`
	Cities      []string `json:"cities"`
}
