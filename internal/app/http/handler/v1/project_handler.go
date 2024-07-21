package v1

import (
	"bas-backend/domain/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
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

func (ph *ProjectHandler) GetProjects(c echo.Context) error {
	partners, err := ph.ProjectUsecase.GetAllProjects(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, partners)
}
