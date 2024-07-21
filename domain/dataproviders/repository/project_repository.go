package repository

import (
	"bas-backend/config"
	"bas-backend/domain/model"
	"context"
	"github.com/go-pg/pg/v10"
	"log"
	"strconv"
)

type ProjectRepository interface {
	// GetAllProjects Получить все проекты
	GetAllProjects(ctx context.Context) ([]model.Project, error)
}

type projectRepository struct {
	db *pg.DB
}

func NewProjectRepository(ctx context.Context, config *config.Config) ProjectRepository {

	connection := pg.Connect(&pg.Options{
		Addr:            config.Database.Host + ":" + strconv.Itoa(config.Database.Port),
		User:            config.Database.User,
		Password:        config.Database.Password,
		Database:        config.Database.Name,
		MaxRetries:      3,
		MaxRetryBackoff: 3,
	})

	err := connection.Ping(ctx)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	return &projectRepository{
		db: connection,
	}
}

// GetAllProjects Получить все проекты
func (p *projectRepository) GetAllProjects(ctx context.Context) ([]model.Project, error) {
	sql := `select id,name,url,status from projects`
	projects := make([]model.Project, 0)

	_, err := p.db.QueryContext(ctx, &projects, sql)
	if err != nil {
		return projects, err
	}

	return projects, nil
}
