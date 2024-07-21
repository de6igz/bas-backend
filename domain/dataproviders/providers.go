package dataproviders

import (
	"bas-backend/config"
	"bas-backend/domain/dataproviders/repository"
	"context"
)

type Providers interface {
	GetPartnerRepository() repository.PartnerRepository
}

type providers struct {
	partnerRepository repository.PartnerRepository
}

func NewProviders(ctx context.Context, cfg *config.Config) Providers {
	partnerRepository := repository.NewPartnerRepository(ctx, cfg)
	return &providers{
		partnerRepository: partnerRepository,
	}
}

func (p *providers) GetPartnerRepository() repository.PartnerRepository {
	return p.partnerRepository
}
