package handler

import (
	"github.com/Futturi/testovoe/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitHandler() http.Handler {
	han := gin.Default()
	api := han.Group("/api")
	{
		api.GET("/:page", h.GetNomers)
		api.DELETE("/", h.Delete)
		api.PUT("/", h.Update)
		api.POST("/", h.Create)
	}
	return han
}
