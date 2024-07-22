package dto

import "bas-backend/domain/model"

type OneProjectDTO struct {
	ShortName     string          `json:"short_name"`
	DeveloperName string          `json:"developer_name"`
	Body          string          `json:"body"`
	Pictures      []model.Picture `json:"pictures"`
	Coordinates   Coordinates     `json:"coordinates"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude" `
	Longitude float64 `json:"longitude"`
}

type ProjectsDto struct {
	Items []ProjectDto `json:"items"`
}

type ProjectDto struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Status string `json:"status"`
}
