package handler

import (
	"net/http"

	"github.com/fprofit/EffectiveMobile/internal/errorResponse"
	"github.com/fprofit/EffectiveMobile/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addUser(c *gin.Context) {
	h.log.Debug("addUser")

	data := models.AddUser{}

	if err := c.BindJSON(&data); err != nil {
		h.log.Debugf("Error parsing request body: %s", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorMsg(http.StatusBadRequest, "Error parsing request body"))
		return
	}

	if err := checkJSONAddUser(data); err != nil {
		h.log.Debugf("Invalid JSON data: %s", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorMsg(http.StatusBadRequest, err.Error()))
		return
	}

	res, err := h.service.AddUser(c, data)
	if err != nil {
		h.log.Errorf("Error adding user: %s", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorMsg(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
	h.log.Infof("User added successfully. ID: %d, Name: %s", res.ID, res.Name)
}
