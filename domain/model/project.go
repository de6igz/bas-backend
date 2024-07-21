package model

import "context"

type Project struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	URL    string `json:"url"`
	Status string `json:"status"`
}

type ProjectRepository interface {
	GetAllProjects(ctx context.Context) ([]Project, error)
}
