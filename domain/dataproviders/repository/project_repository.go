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
	GetProjectById(ctx context.Context, id int) (model.Project, error)
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

func (p *projectRepository) GetProjectById(ctx context.Context, id int) (model.Project, error) {
	pictures := make([]model.Picture, 0, 10)
	picturesSql := `select url from pictures where pictures.project_id = ?`
	_, err := p.db.QueryContext(ctx, &pictures, picturesSql, id)
	if err != nil {
		log.Printf("error getting pictures: %v", err)
		return model.Project{}, err
	}

	project := model.Project{
		Pictures: pictures,
	}
	projectsSql := `select  project_name,
					builder_name,
					body,
					coordinates[0] as latitude, 
					coordinates[1] as longitude
			from projects
			where id = ?`

	_, err = p.db.QueryContext(ctx, &project, projectsSql, id)
	if err != nil {
		return project, err
	}

	return project, nil
}
