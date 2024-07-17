package model

type Partner struct {
	URL         string `json:"url"`
	Description string `json:"description"`
}

type PartnerRepository interface {
	GetAll() ([]*Partner, error)
}
