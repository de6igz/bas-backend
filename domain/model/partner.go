package model

import "context"

type Partner struct {
	URL         string `json:"url"`
	Description string `json:"description"`
}

type PartnerRepository interface {
	GetAllPartners(ctx context.Context) ([]Partner, error)
}
