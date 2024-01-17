package handler

import (
	"github.com/fprofit/EffectiveMobile/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type service interface {
	AddUser(c *gin.Context, addUser models.AddUser) (models.ResponseUser, error)
	DelUser(c *gin.Context, id int64) error
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
	router.POST("/user", h.addUser)
	router.GET("/user", h.getUsers)
	router.PUT("/user/:id", h.updUser)
	router.DELETE("/user/:id", h.delUser)
	return router
}
