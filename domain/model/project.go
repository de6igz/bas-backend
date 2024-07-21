package model

import "context"

type Project struct {
	ID          int       `json:"id" pg:"id"`
	Name        string    `json:"name" pg:"name"`
	URL         string    `json:"url" pg:"url"`
	Status      string    `json:"status" pg:"status"`
	ShortName   string    `json:"short_name" pg:"project_name"`
	BuilderName string    `json:"builder_name" pg:"builder_name"`
	Body        string    `json:"body" pg:"body"`
	Pictures    []Picture `json:"pictures" pg:"pictures"`
	Latitude    float64   `json:"latitude" pg:"latitude"`
	Longitude   float64   `json:"longitude" pg:"longitude"`
}

type Picture struct {
	Url string `json:"url" pg:"url"`
}

type ProjectRepository interface {
	GetAllProjects(ctx context.Context) ([]Project, error)
}
