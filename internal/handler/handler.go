package handler

import (
	"github.com/fprofit/EffectiveMobile/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type service interface {
	AddPerson(addUser models.Person) (models.EnrichedPerson, error)
	DelPerson(id int64) error
	UpdPerson(person models.EnrichedPerson) (models.EnrichedPerson, error)
	GetPersonsByFilter(filter models.PersonFilter) (models.PersonList, error)
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
	router.POST("/person", h.addPerson)
	router.GET("/persons", h.getPersons)
	router.PUT("/person/:id", h.updPerson)
	router.DELETE("/person/:id", h.delPerson)
	return router
}
