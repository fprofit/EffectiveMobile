package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fprofit/EffectiveMobile/internal/errorResponse"

	"github.com/gin-gonic/gin"
)

// @Summary Удаление по идентификатору
// @Description Удаляет человека по указанному идентификатору.
// @ID delete-person-by-id
// @Param id path int true "Идентификатор человека" Format(int64)
// @Produce json
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} errorResponse.ErrorResponse
// @Failure 404 {object} errorResponse.ErrorResponse
// @Failure 500 {object} errorResponse.ErrorResponse
// @Router /person/{id} [delete]
func (h *Handler) delPerson(c *gin.Context) {
	h.log.Debug("Handler delPerson")

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.log.Debugf("Error parsing person ID: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorMsg(http.StatusBadRequest, "Error parsing user ID"))
		return
	}
	h.log.Debugf("Handler delPerson delete person ID: %d", id)
	if err := h.service.DelPerson(id); err != nil {
		h.log.Errorf("Error deleting person: %s", err.Error())
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse.NewErrorMsg(http.StatusNotFound, fmt.Sprintf("Not foun ID: %d", id)))
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse.ErrorStatusInternalServerError)
		}
		return
	}
	h.log.Debugf("Handler delPerson delete person successfully ID: %d", id)
	c.JSON(http.StatusOK, SuccessResponse{Message: fmt.Sprintf("Person with ID %d deleted successfully", id)})
}
