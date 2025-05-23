package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Define algunas métricas básicas
	authRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "auth_requests_total",
		Help: "El número total de solicitudes al servicio de autenticación",
	})

	authUsers = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "auth_users_active",
		Help: "El número actual de usuarios activos",
	})
)

// SetupPrometheusMetrics inicializa las métricas y registra algunos valores de ejemplo
func SetupPrometheusMetrics() {
	// Aumenta la métrica de solicitudes totales
	authRequests.Inc()

	// Establece un valor de ejemplo para usuarios activos
	authUsers.Set(10)
}

// Uso normal con JSON
func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Endpoint que expone las métricas de Prometheus
func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
