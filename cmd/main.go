package main

import (
	"bas-backend/config"
	"bas-backend/internal/app/http"
	"bas-backend/pkg/echo"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	e := echo.New()
	http.RegisterRoutes(e, cfg)
	address := fmt.Sprintf(":%d", cfg.Server.Port)
	e.Logger.Fatal(e.Start(address))
}
