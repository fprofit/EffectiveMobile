package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fprofit/EffectiveMobile/internal/errorResponse"
	"github.com/gin-gonic/gin"
)

func (h *Handler) delUser(c *gin.Context) {
	h.log.Debug("delUser")

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.log.Debugf("Error parsing user ID: %s", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorMsg(http.StatusBadRequest, "Error parsing user ID"))
		return
	}

	if err := h.service.DelUser(c, id); err != nil {
		h.log.Errorf("Error deleting user: %s", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorMsg(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User with ID %d deleted successfully", id)})
	h.log.Infof("User with ID %d deleted successfully", id)
}
