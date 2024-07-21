package dataproviders

import (
	"bas-backend/config"
	"bas-backend/domain/dataproviders/repository"
	"context"
)

type Providers interface {
	GetPartnerRepository() repository.PartnerRepository
	GetProjectRepository() repository.ProjectRepository
	GetDocumentRepository() repository.DocRepository
}

type providers struct {
	partnerRepository  repository.PartnerRepository
	projectRepository  repository.ProjectRepository
	documentRepository repository.DocRepository
}

func NewProviders(ctx context.Context, cfg *config.Config) Providers {
	partnerRepository := repository.NewPartnerRepository(ctx, cfg)
	projectRepository := repository.NewProjectRepository(ctx, cfg)
	documentRepository := repository.NewDocRepository(ctx, cfg)
	return &providers{
		partnerRepository:  partnerRepository,
		projectRepository:  projectRepository,
		documentRepository: documentRepository,
	}
}

func (p *providers) GetPartnerRepository() repository.PartnerRepository {
	return p.partnerRepository
}

func (p *providers) GetProjectRepository() repository.ProjectRepository {
	return p.projectRepository
}

func (p *providers) GetDocumentRepository() repository.DocRepository {
	return p.documentRepository
}
