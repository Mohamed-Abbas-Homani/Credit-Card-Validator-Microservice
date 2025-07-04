package middleware

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	requestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "card_validation_requests_total",
			Help: "Total number of card validation requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	requestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "card_validation_duration_seconds",
			Help: "Duration of card validation requests",
		},
		[]string{"method", "endpoint"},
	)

	validationErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "card_validation_errors_total",
			Help: "Total number of card validation errors",
		},
		[]string{"error_type"},
	)
)

func RequestID() echo.MiddlewareFunc {
	return middleware.RequestID()
}

func Metrics() echo.MiddlewareFunc {
	return echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			duration := time.Since(start)
			status := c.Response().Status

			requestsTotal.WithLabelValues(
				c.Request().Method,
				c.Path(),
				strconv.Itoa(status),
			).Inc()

			requestDuration.WithLabelValues(
				c.Request().Method,
				c.Path(),
			).Observe(duration.Seconds())

			if status >= 400 {
				validationErrors.WithLabelValues("http_error").Inc()
			}

			return err
		}
	})
}
