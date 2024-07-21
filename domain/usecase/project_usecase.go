package usecase

import (
	"bas-backend/domain/dataproviders"
	"bas-backend/domain/dataproviders/repository"
	"bas-backend/domain/model"
	"context"
)

type ProjectUsecase interface {
	GetAllProjects(ctx context.Context) ([]model.Project, error)
}

type projectUsecase struct {
	projectRepo repository.ProjectRepository
}

func NewProjectUsecase(providers dataproviders.Providers) ProjectUsecase {

	return &projectUsecase{
		projectRepo: providers.GetProjectRepository(),
	}
}

func (p *projectUsecase) GetAllProjects(ctx context.Context) ([]model.Project, error) {
	items, err := p.projectRepo.GetAllProjects(ctx)
	if err != nil {

		return nil, err
	}

	return items, nil
}
