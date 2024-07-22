package metrics

import (
	"errors"
	"fmt"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method", "status"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method", "status"},
	)
)

func Init(e *echo.Echo) {

	e.Use(echoprometheus.NewMiddleware("echo")) // adds middleware to gather metrics
}

func StartMetricsServer(port int) {
	e := echo.New()
	e.GET("/metrics", echoprometheus.NewHandler()) // adds route to serve gathered metrics

	address := fmt.Sprintf(":%d", port)
	go func() {
		if err := e.Start(address); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting metrics server: %v", err)
		}
	}()
}
