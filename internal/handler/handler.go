package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type service interface {
}

type Handler struct {
	service service
	log     *logrus.Logger
}

func NewHandler(service service, log *logrus.Logger) *Handler {
	log.Debug("Initializing Handlers...")
	return &Handler{service: service, log: log}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/user", h.getUsers)
	router.DELETE("/user/:id", h.delUser)
	router.PUT("/user/:id", h.updUser)
	router.POST("/user", h.addUser)
	return router
}
