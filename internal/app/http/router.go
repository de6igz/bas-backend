package http

import (
	"bas-backend/config"
	"bas-backend/domain/dataproviders"
	"bas-backend/domain/usecase"
	"bas-backend/internal/app/http/handler"
	"context"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(ctx context.Context, e *echo.Echo, cfg *config.Config) {
	// Инициализация датапровайдеров
	providers := dataproviders.NewProviders(ctx, cfg)

	// Инициализация юзкейсов
	partnerUsecase := usecase.NewPartnerUsecase(providers)

	// Инициализация ручек
	handler.NewPartnerHandler(e, "/v1/partners", partnerUsecase)

}
