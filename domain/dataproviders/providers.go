package dataproviders

import (
	"bas-backend/config"
	"bas-backend/domain/dataproviders/repository"
	"context"
)

type Providers interface {
	GetPartnerRepository() repository.PartnerRepository
	GetProjectRepository() repository.ProjectRepository
}

type providers struct {
	partnerRepository repository.PartnerRepository
	projectRepository repository.ProjectRepository
}

func NewProviders(ctx context.Context, cfg *config.Config) Providers {
	partnerRepository := repository.NewPartnerRepository(ctx, cfg)
	projectRepository := repository.NewProjectRepository(ctx, cfg)
	return &providers{
		partnerRepository: partnerRepository,
		projectRepository: projectRepository,
	}
}

func (p *providers) GetPartnerRepository() repository.PartnerRepository {
	return p.partnerRepository
}

func (p *providers) GetProjectRepository() repository.ProjectRepository {
	return p.projectRepository
}
