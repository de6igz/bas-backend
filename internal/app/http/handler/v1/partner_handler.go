package v1

import (
	"bas-backend/domain/model"
	"bas-backend/domain/usecase"
	"bas-backend/internal/app/http/handler/v1/dto"
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

// GetPartners  получить список всех партнеров
// @Summary получить список всех партнеров
// @Success      200 {object}  dto.PartnersDto
// @Router       /v1/partners [get]
func (ph *PartnerHandler) GetPartners(c echo.Context) error {
	partners, err := ph.PartnerUsecase.Fetch(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ph.prepareResponse(partners))
}

func (ph *PartnerHandler) prepareResponse(data []model.Partner) dto.PartnersDto {
	items := make([]dto.Partner, 0, len(data))
	for _, item := range data {
		items = append(items, dto.Partner{
			URL:         item.URL,
			Description: item.Description,
		})
	}

	return dto.PartnersDto{
		Items: items,
	}
}
