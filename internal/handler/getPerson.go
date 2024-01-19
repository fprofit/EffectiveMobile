package handler

import (
	"net/http"

	"github.com/fprofit/EffectiveMobile/internal/errorResponse"
	"github.com/fprofit/EffectiveMobile/internal/models"
	"github.com/fprofit/EffectiveMobile/internal/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getPersons(c *gin.Context) {
	h.log.Debug("Handler getPersons")

	var filter models.PersonFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		h.log.Debugf("Error parsing person filter: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorMsg(http.StatusBadRequest, "Invalid filter parameters"))
		return
	}
	if filter.Sort == nil {
		id := "id"
		filter.Sort = &id
	}
	if err := checkFilterGetPerson(filter); err != nil {
		h.log.Debugf("Error checkFilterGetPerson: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorMsg(http.StatusBadRequest, err.Error()))
		return
	}
	h.log.Debugf("Handler getPersons filter: %s", utils.StructToString(filter))
	persons, err := h.service.GetPersonsByFilter(filter)
	if err != nil {
		h.log.Errorf("Error getPersons person: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse.ErrorStatusInternalServerError)
		return
	}

	h.log.Debugf("Handler getPersons res data: %s", utils.StructToString(persons))
	c.JSON(http.StatusOK, persons)
}
