package handler

import (
	"net/http"

	"github.com/Solutions-Corp/chetapp-backend/fleet-management/internal/model"
	"github.com/Solutions-Corp/chetapp-backend/fleet-management/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BusHandler struct {
	busService service.BusService
}

func NewBusHandler(busService service.BusService) *BusHandler {
	return &BusHandler{
		busService: busService,
	}
}

func (h *BusHandler) CreateBus(c *gin.Context) {
	var bus model.Bus
	if err := c.ShouldBindJSON(&bus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	bus.CreatedBy = userID.(uuid.UUID)
	bus.UpdatedBy = userID.(uuid.UUID)

	if bus.GpsID != nil && *bus.GpsID == uuid.Nil {
		bus.GpsID = nil
	}

	if err := h.busService.CreateBus(&bus); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, bus)
}

func (h *BusHandler) GetBusByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	bus, err := h.busService.GetBusByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bus no encontrado"})
		return
	}

	c.JSON(http.StatusOK, bus)
}

func (h *BusHandler) GetAllBuses(c *gin.Context) {
	buses, err := h.busService.GetAllBuses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"buses": buses})
}
