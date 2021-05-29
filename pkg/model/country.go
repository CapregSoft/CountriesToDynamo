package model

// Country store the model for the country
type Country struct {
	CountryName string   `json:"name"`
	Cities      []string `json:"cities"`
}
