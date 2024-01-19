package handler

import (
	"net/http"

	"github.com/fprofit/EffectiveMobile/internal/errorResponse"
	"github.com/fprofit/EffectiveMobile/internal/models"
	"github.com/fprofit/EffectiveMobile/internal/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Получение данных с различными фильтрами, сортировкой и пагинацией
// @Description Возвращает список людей в соответствии с указанными фильтрами, сортировкой и пагинацией.
// @ID get-persons
// @Accept json
// @Produce json
// @Param limit query int false "Ограничение количества записей в результате"
// @Param offset query int false "Смещение относительно начала списка"
// @Param age query int false "Фильтрация по возрасту (age > 0)"
// @Param min_age query int false "Минимальный возраст"
// @Param max_age query int false "Максимальный возраст"
// @Param gender query string false "Фильтрация по полу ('male' или 'female')"
// @Param name query string false "Фильтрация по имени"
// @Param surname query string false "Фильтрация по фамилии"
// @Param country query string false "Фильтрация по стране (Формат ISO 3166-1 alpha-2)"
// @Param sort query string false "Поле для сортировки (id, name, surname, patronymic, age, gender, country_id)"
// @Success 200 {object} models.PersonList "Список людей с информацией о пагинации"
// @Failure 400 {object} errorResponse.ErrorResponse "Некорректный возраст, пол, или id страны"
// @Failure 500 {object} errorResponse.ErrorResponse "База данных не отвечает"
// @Router /persons [get]
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
