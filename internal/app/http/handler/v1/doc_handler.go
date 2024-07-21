package v1

import (
	"bas-backend/domain/usecase"
	"bas-backend/internal/app/http/handler/v1/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DocumentHandler struct {
	documentUsecase usecase.DocumentsUsecase
}

func NewDocumentHandler(e *echo.Echo, route string, pu usecase.DocumentsUsecase) {
	handler := &DocumentHandler{
		documentUsecase: pu,
	}
	e.GET(route, handler.GetDocuments)
}

func (ph *DocumentHandler) GetDocuments(c echo.Context) error {
	docs, err := ph.documentUsecase.GetAllDocuments(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, dto.DocDto{Items: docs})
}
