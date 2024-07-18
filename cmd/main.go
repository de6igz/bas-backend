package main

import (
	"bas-backend/config"
	"bas-backend/internal/app/http"
	"bas-backend/pkg/echo"
	"context"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}
	ctx := context.Background()
	e := echo.New()
	http.RegisterRoutes(ctx, e, cfg)
	address := fmt.Sprintf(":%d", cfg.Server.Port)
	e.Logger.Fatal(e.Start(address))
}
