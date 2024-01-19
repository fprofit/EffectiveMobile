package models

type Person struct {
	Name       *string `json:"name"`
	Surname    *string `json:"surname"`
	Patronymic *string `json:"patronymic,omitempty"`
}

type EnrichedPerson struct {
	ID         int64   `json:"id" db:"id"`
	Name       *string `json:"name" db:"name"`
	Surname    *string `json:"surname" db:"surname"`
	Patronymic *string `json:"patronymic,omitempty" db:"patronymic"`
	Age        *int    `json:"age" db:"age"`
	Gender     *string `json:"gender" db:"gender"`
	CountryID  *string `json:"country_id" db:"country_id"`
}

type PersonList struct {
	Persons []EnrichedPerson `json:"persons"`
	Offset  int64            `json:"offset"`
	Limit   int64            `json:"limit"`
	Count   int64            `json:"count"`
}

type PersonFilter struct {
	Limit   int64   `form:"limit"`
	Offset  int64   `form:"offset"`
	Sort    *string `form:"sort"`
	MinAge  *int    `form:"min_age"`
	MaxAge  *int    `form:"max_age"`
	Age     *int    `form:"age"`
	Name    *string `form:"name"`
	Surname *string `form:"surname"`
	Gender  *string `form:"gender"`
	Country *string `form:"country"`
}
