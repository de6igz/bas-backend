package http

import (
	"bas-backend/config"
	"bas-backend/domain/dataproviders"
	"bas-backend/domain/usecase"
	"bas-backend/internal/app/http/handler/v1"
	"context"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(ctx context.Context, e *echo.Echo, cfg *config.Config) {
	// Инициализация датапровайдеров
	providers := dataproviders.NewProviders(ctx, cfg)

	// Инициализация юзкейсов
	partnerUsecase := usecase.NewPartnerUsecase(providers)
	projectUsecase := usecase.NewProjectUsecase(providers)
	documentUsecase := usecase.NewDocumentUsecase(providers)

	// Инициализация ручек
	v1.NewPartnerHandler(e, "/v1/partners", partnerUsecase)
	v1.NewProjectHandler(e, "/v1/projects", projectUsecase)
	v1.NewDocumentHandler(e, "/v1/docs", documentUsecase)

}
