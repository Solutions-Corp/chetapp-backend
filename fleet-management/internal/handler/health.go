package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Define algunas métricas básicas para fleet-management
	fleetBuses = promauto.NewCounter(prometheus.CounterOpts{
		Name: "fleet_buses_total",
		Help: "El número total de autobuses registrados",
	})

	fleetGpsDevices = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "fleet_gps_devices_active",
		Help: "El número actual de dispositivos GPS activos",
	})

	fleetCompanies = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "fleet_companies_total",
		Help: "El número total de compañías registradas",
	})
)

// SetupPrometheusMetrics inicializa las métricas y registra algunos valores de ejemplo
func SetupPrometheusMetrics() {
	// Establece valores de ejemplo
	fleetBuses.Inc()
	fleetGpsDevices.Set(5)
	fleetCompanies.Set(2)
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
