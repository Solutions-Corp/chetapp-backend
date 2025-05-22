package handler

import (
	"net/http"

	"github.com/Solutions-Corp/chetapp-backend/fleet-management/internal/model"
	"github.com/Solutions-Corp/chetapp-backend/fleet-management/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GpsHandler struct {
	gpsService service.GpsService
}

func NewGpsHandler(gpsService service.GpsService) *GpsHandler {
	return &GpsHandler{
		gpsService: gpsService,
	}
}

func (h *GpsHandler) CreateGps(c *gin.Context) {
	var gps model.Gps
	if err := c.ShouldBindJSON(&gps); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	gps.CreatedBy = userID.(uuid.UUID)
	gps.UpdatedBy = userID.(uuid.UUID)

	if err := h.gpsService.CreateGps(&gps); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gps)
}

func (h *GpsHandler) GetGpsByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	gps, err := h.gpsService.GetGpsByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "GPS no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gps)
}

func (h *GpsHandler) GetAllGps(c *gin.Context) {
	gpss, err := h.gpsService.GetAllGps()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"gps": gpss})
}
