package repository

import (
	"bas-backend/domain/model"
	"context"
	"database/sql"
	"log"
)

type ProjectRepository interface {
	// GetAllProjects Получить все проекты
	GetAllProjects(ctx context.Context) ([]model.Project, error)
	GetProjectById(ctx context.Context, id int) (model.Project, error)
}

type projectRepository struct {
	db *sql.DB
}

func NewProjectRepository(ctx context.Context, db *sql.DB) ProjectRepository {
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	return &projectRepository{
		db: db,
	}
}

// GetAllProjects Получить все проекты
func (p *projectRepository) GetAllProjects(ctx context.Context) ([]model.Project, error) {
	sqlText := `select id, full_name, url, status from projects`
	rows, err := p.db.QueryContext(ctx, sqlText)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := make([]model.Project, 0)
	for rows.Next() {
		var project model.Project
		if err := rows.Scan(&project.ID, &project.Name, &project.URL, &project.Status); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, rows.Err()
}

func (p *projectRepository) GetProjectById(ctx context.Context, id int) (model.Project, error) {
	pictures := make([]model.Picture, 0, 10)
	picturesSql := `select url from pictures where pictures.project_id = ?`

	rows, err := p.db.QueryContext(ctx, picturesSql, id)
	if err != nil {
		log.Printf("error getting pictures: %v", err)
		return model.Project{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var picture model.Picture
		if err := rows.Scan(&picture.Url); err != nil {
			return model.Project{}, err
		}
		pictures = append(pictures, picture)
	}
	if err := rows.Err(); err != nil {
		return model.Project{}, err
	}

	project := model.Project{
		Pictures: pictures,
	}
	projectsSql := `select  short_name,
					builder_name,
					body,
					latitude, 
					longitude
			from projects
			where id = ?`

	err = p.db.QueryRowContext(ctx, projectsSql, id).Scan(&project.ShortName, &project.BuilderName, &project.Body, &project.Latitude, &project.Longitude)
	if err != nil {
		return project, err
	}

	return project, nil
}
