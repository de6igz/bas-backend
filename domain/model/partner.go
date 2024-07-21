package model

import "context"

type Partner struct {
	URL         string `pg:"url"`
	Description string `pg:"description"`
}

type PartnerRepository interface {
	GetAllPartners(ctx context.Context) ([]Partner, error)
}
