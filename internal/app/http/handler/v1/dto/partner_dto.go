package dto

type Partner struct {
	URL         string `json:"url"`
	Description string `json:"description"`
}

type PartnersDto struct {
	Items []Partner `json:"items"`
}
