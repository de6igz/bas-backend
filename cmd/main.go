package main

import (
	"bas-backend/internal/app/http"
	"bas-backend/pkg/echo"
)

func main() {
	e := echo.New()
	http.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
