package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Define algunas métricas básicas para routes
	routesTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "routes_total",
		Help: "El número total de rutas registradas",
	})

	routesActive = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "routes_active",
		Help: "El número de rutas activas",
	})

	wsConnections = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "ws_connections_active",
		Help: "El número de conexiones WebSocket activas",
	})
)

// SetupPrometheusMetrics inicializa las métricas y registra algunos valores de ejemplo
func SetupPrometheusMetrics() {
	// Establece valores de ejemplo
	routesTotal.Inc()
	routesActive.Set(1)
	wsConnections.Set(0)
}

// Endpoint de salud tradicional
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
