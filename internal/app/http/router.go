package http

import (
	"bas-backend/config"
	"bas-backend/domain/repository"
	"bas-backend/domain/usecase"
	"bas-backend/internal/app/http/handler"
	"context"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(ctx context.Context, e *echo.Echo, cfg *config.Config) {
	partnerRepo := repository.NewPartnerRepository(ctx, cfg)
	partnerUsecase := usecase.NewPartnerUsecase(partnerRepo)
	handler.NewPartnerHandler(e, partnerUsecase)

	// Повторить для проектов и документов
}
