package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fprofit/EffectiveMobile/internal/errorResponse"
	"github.com/fprofit/EffectiveMobile/internal/models"
	"github.com/fprofit/EffectiveMobile/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Изменение сущности по идентификатору
// @Description Изменяет данные о человеке с указанным идентификатором
// @ID update-person
// @Accept json
// @Produce json
// @Param id path int true "Идентификатор человека"
// @Param request_body body models.EnrichedPerson true "Данные для обновления человека"
// @Success 200 {object} models.EnrichedPerson "Данные о человеке после обновления"
// @Failure 400 {object} errorResponse.ErrorResponse "Некорректное id, тело запроса, возраст, пол или id страны"
// @Failure 404 {object} errorResponse.ErrorResponse "id не найдено"
// @Failure 500 {object} errorResponse.ErrorResponse "База данных не отвечает"
// @Router /person/{id} [put]
func (h *Handler) updPerson(c *gin.Context) {
	h.log.Debug("Handler updPerson")

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.log.Debugf("Error parsing person ID: %s", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorMsg(http.StatusBadRequest, "Error parsing user ID"))
		return
	}

	data := models.EnrichedPerson{}

	if err := c.BindJSON(&data); err != nil {
		h.log.Debugf("Error parsing request body: %s", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorMsg(http.StatusBadRequest, "Error parsing request body"))
		return
	}

	data.ID = id
	if err := checkJSONUpdPerson(data); err != nil {
		h.log.Debugf("Error checkJSONUpdPerson: %s", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse.NewErrorMsg(http.StatusBadRequest, err.Error()))
		return
	}

	h.log.Debugf("Handler updPerson data: %s", utils.StructToString(data))
	res, err := h.service.UpdPerson(data)
	if err != nil {
		h.log.Errorf("Error update person: %s", err)
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse.NewErrorMsg(http.StatusNotFound, fmt.Sprintf("Not foun ID: %d", id)))
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse.ErrorStatusInternalServerError)
		}
		return
	}

	h.log.Debugf("Handler updPerson update data: %s", utils.StructToString(res))
	c.JSON(http.StatusOK, res)
}
