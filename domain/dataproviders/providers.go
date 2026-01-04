package dataproviders

import (
	"bas-backend/config"
	"bas-backend/domain/dataproviders/repository"
	"context"
	"log"
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
	db, err := initSQLite(ctx, cfg)
	if err != nil {
		log.Fatalf("init sqlite: %v", err)
	}

	partnerRepository := repository.NewPartnerRepository(ctx, db)
	projectRepository := repository.NewProjectRepository(ctx, db)
	documentRepository := repository.NewDocRepository(ctx, db)
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
