package v1

import (
	"bas-backend/domain/model"
	"bas-backend/domain/usecase"
	"bas-backend/internal/app/http/handler/v1/dto"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type ProjectHandler struct {
	ProjectUsecase usecase.ProjectUsecase
}

func NewProjectHandler(e *echo.Echo, route string, pu usecase.ProjectUsecase) {
	handler := &ProjectHandler{
		ProjectUsecase: pu,
	}
	e.GET(route, handler.GetProjects)

}

// GetProjects получить список всех проектов
// @Summary получить список всех проектов. Чтобы получить инфу об одном проекте нужно указать query param
// @Success      200 {object}  dto.ProjectsDto
// @Param        id   query      int  false  "ID проекта"
// @Router       /v1/projects [get]
func (ph *ProjectHandler) GetProjects(c echo.Context) error {
	idStr := c.QueryParam("id")

	if idStr != "" {
		response, err := ph.getProjectByID(c.Request().Context(), idStr)
		if err != nil {
			log.Printf("get project err %v", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, response)

	}

	partners, err := ph.ProjectUsecase.GetAllProjects(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ph.prepareProjects(partners))
}

func (ph *ProjectHandler) getProjectByID(ctx context.Context, idStr string) (dto.OneProjectDTO, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return dto.OneProjectDTO{}, errors.New("invalid project id")
	}
	project, err := ph.ProjectUsecase.GetProjectByID(ctx, id)

	return dto.OneProjectDTO{
		ShortName:     project.ShortName,
		DeveloperName: project.BuilderName,
		Body:          project.Body,
		Pictures:      project.Pictures,
		Coordinates: dto.Coordinates{
			Longitude: project.Longitude,
			Latitude:  project.Latitude,
		},
	}, nil

}

func (ph *ProjectHandler) prepareProjects(data []model.Project) dto.ProjectsDto {
	projects := make([]dto.ProjectDto, 0, len(data))

	for _, project := range data {
		projects = append(projects, dto.ProjectDto{
			Id:     project.ID,
			Name:   project.Name,
			Url:    project.URL,
			Status: project.Status,
		})
	}

	return dto.ProjectsDto{
		Items: projects,
	}
}
