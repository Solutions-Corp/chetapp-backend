package handler

import (
	"net/http"

	"github.com/Solutions-Corp/chetapp-backend/routes/internal/model"
	"github.com/Solutions-Corp/chetapp-backend/routes/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RouteHandler struct {
	routeService service.RouteService
}

func NewRouteHandler(routeService service.RouteService) *RouteHandler {
	return &RouteHandler{
		routeService: routeService,
	}
}

func (h *RouteHandler) RegisterRoutes(router gin.IRouter) {
	routes := router.Group("/routes")
	{
		routes.POST("", h.CreateRoute)
		routes.GET("", h.GetAllRoutes)
		routes.GET("/:id", h.GetRouteByID)
		routes.PUT("/:id", h.UpdateRoute)
		routes.DELETE("/:id", h.DeleteRoute)
		routes.POST("/upload-gpx", h.UploadGPX)
	}
}

func (h *RouteHandler) CreateRoute(c *gin.Context) {
	var route model.Route
	if err := c.ShouldBindJSON(&route); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	user, _ := uuid.Parse(userID.(string))
	route.CreatedBy = user
	route.UpdatedBy = user

	if err := h.routeService.CreateRoute(&route); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, route)
}

func (h *RouteHandler) GetAllRoutes(c *gin.Context) {
	routes, err := h.routeService.GetAllRoutes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, routes)
}

func (h *RouteHandler) GetRouteByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	route, err := h.routeService.GetRouteByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ruta no encontrada"})
		return
	}

	c.JSON(http.StatusOK, route)
}

func (h *RouteHandler) UpdateRoute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	existingRoute, err := h.routeService.GetRouteByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ruta no encontrada"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	user, _ := uuid.Parse(userID.(string))

	var updatedRoute model.Route
	if err := c.ShouldBindJSON(&updatedRoute); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedRoute.ID = existingRoute.ID
	updatedRoute.CreatedBy = existingRoute.CreatedBy
	updatedRoute.UpdatedBy = user

	if err := h.routeService.UpdateRoute(&updatedRoute); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedRoute)
}

// DeleteRoute elimina una ruta por su ID
func (h *RouteHandler) DeleteRoute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.routeService.DeleteRoute(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ruta eliminada correctamente"})
}

func (h *RouteHandler) UploadGPX(c *gin.Context) {
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}
	userID, ok := userIDValue.(uuid.UUID)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID in context"})
		return
	}

	file, header, err := c.Request.FormFile("gpx_file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al recibir el archivo: " + err.Error()})
		return
	}
	defer file.Close()

	name := c.PostForm("name")
	if name == "" {
		name = header.Filename
	}

	route, err := h.routeService.ProcessGPXFile(file, name, userID, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el archivo GPX: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, route)
}
