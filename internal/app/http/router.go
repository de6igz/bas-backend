package http

import (
	"bas-backend/domain/repository"
	"bas-backend/domain/usecase"
	"bas-backend/internal/app/http/handler"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	partnerRepo := repository.NewPartnerRepository()
	partnerUsecase := usecase.NewPartnerUsecase(partnerRepo)
	handler.NewPartnerHandler(e, partnerUsecase)

	// Повторить для проектов и документов
}
