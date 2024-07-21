package usecase

import (
	"bas-backend/domain/dataproviders"
	"bas-backend/domain/dataproviders/repository"
	"bas-backend/domain/model"
	"context"
)

type ProjectUsecase interface {
	GetAllProjects(ctx context.Context) ([]model.Project, error)
	GetProjectByID(ctx context.Context, id int) (model.Project, error)
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

func (p *projectUsecase) GetProjectByID(ctx context.Context, id int) (model.Project, error) {
	project, err := p.projectRepo.GetProjectById(ctx, id)
	if err != nil {
		return model.Project{}, err
	}

	return project, nil
}
