package repository

import (
	"fmt"
	"strings"

	"github.com/fprofit/EffectiveMobile/internal/models"
)

func createFilterQuery(filter models.PersonFilter) string {
	var conditions []string

	if filter.MinAge != nil {
		conditions = append(conditions, fmt.Sprintf("age >= %d", *filter.MinAge))
	}
	if filter.MaxAge != nil {
		conditions = append(conditions, fmt.Sprintf("age <= %d", *filter.MaxAge))
	}
	if filter.Gender != nil {
		conditions = append(conditions, fmt.Sprintf("gender = '%s'", *filter.Gender))
	}
	if filter.Name != nil {
		conditions = append(conditions, fmt.Sprintf("name = '%s'", *filter.Name))
	}
	if filter.Age != nil {
		conditions = append(conditions, fmt.Sprintf("age = %d", *filter.Age))
	}
	if filter.Surname != nil {
		conditions = append(conditions, fmt.Sprintf("surname = '%s'", *filter.Surname))
	}
	if filter.Country != nil {
		conditions = append(conditions, fmt.Sprintf("country_id = '%s'", *filter.Country))
	}

	filterQuery := strings.Join(conditions, " AND ")

	if len(conditions) > 0 {
		filterQuery = fmt.Sprintf("WHERE %s", filterQuery)
	}

	return filterQuery
}
