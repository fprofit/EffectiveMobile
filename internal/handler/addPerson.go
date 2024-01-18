package handler

import (
	"net/http"

	"github.com/fprofit/EffectiveMobile/internal/errorResponse"
	"github.com/fprofit/EffectiveMobile/internal/models"
	"github.com/fprofit/EffectiveMobile/internal/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addPerson(c *gin.Context) {
	h.log.Debug("Handler addPerson")

	data := models.Person{}

	if err := c.BindJSON(&data); err != nil {
		h.log.Debugf("Error parsing request body: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorMsg(http.StatusBadRequest, "Error parsing request body"))
		return
	}

	if err := checkJSONAddPerson(data); err != nil {
		h.log.Debugf("Invalid JSON data: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorMsg(http.StatusBadRequest, err.Error()))
		return
	}

	h.log.Debugf("Handler addPerson get data: %s", utils.StructToString(data))

	res, err := h.service.AddPerson(data)
	if err != nil {
		h.log.Errorf("Error adding person: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse.ErrorStatusInternalServerError)
		return
	}

	h.log.Debugf("Handler addPerson res data: %s", utils.StructToString(res))

	c.JSON(http.StatusOK, res)
}
