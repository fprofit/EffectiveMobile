package models

type AddUser struct {
	Name       *string `json:"name"`
	Surname    *string `json:"surname"`
	Patronymic *string `json:"patronymic,omitempty"`
}

type ResponseUser struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Surname    string `json:"surname" db:"surname"`
	Patronymic string `json:"patronymic" db:"patronymic"`
	Age        int    `db:"age" db:"age"`
	Gender     string `db:"gender" db:"gender"`
	CountryID  string `db:"country_id" db:"country_id"`
}
