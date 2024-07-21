package handler

import (
	"bas-backend/domain/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PartnerHandler struct {
	PartnerUsecase usecase.PartnerUsecase
}

func NewPartnerHandler(e *echo.Echo, route string, pu usecase.PartnerUsecase) {
	handler := &PartnerHandler{
		PartnerUsecase: pu,
	}
	e.GET(route, handler.GetPartners)
}

func (ph *PartnerHandler) GetPartners(c echo.Context) error {
	partners, err := ph.PartnerUsecase.Fetch(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, partners)
}
