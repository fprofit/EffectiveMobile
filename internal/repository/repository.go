package repository

import (
	"database/sql"
	"fmt"

	"github.com/fprofit/EffectiveMobile/internal/models"
	"github.com/fprofit/EffectiveMobile/internal/utils"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewRepository(db *sqlx.DB, log *logrus.Logger) *Repository {
	return &Repository{
		db:  db,
		log: log,
	}
}

func (r *Repository) GetPersonByID(id int64) (models.EnrichedPerson, error) {
	r.log.Debugf("Repository GetPersonByID")

	var person models.EnrichedPerson
	r.log.Debugf("Fetching person data for ID: %d", id)

	query := "SELECT * FROM persons WHERE id = $1"
	err := r.db.Get(&person, query, id)
	if err != nil {
		r.log.Errorf("Error fetching person data for ID %d: %v", id, err)
		return person, err
	}
	r.log.Infof("Data received successfully. %s", utils.StructToString(person))
	return person, nil
}

func (r *Repository) UpdPerson(person models.EnrichedPerson) (models.EnrichedPerson, error) {
	r.log.Debug("Repository UpdPerson")

	query := `
		UPDATE persons 
		SET name=:name, surname=:surname, patronymic=:patronymic, age=:age, gender=:gender, country_id=:country_id
		WHERE id=:id
		RETURNING id, name, surname, patronymic, age, gender, country_id
	`
	result, err := r.db.NamedQuery(query, person)
	if err != nil {
		r.log.Errorf("Error executing SQL query: %s", err)
		return models.EnrichedPerson{}, err
	}
	defer result.Close()

	if result.Next() {
		if err := result.StructScan(&person); err != nil {
			r.log.Errorf("Error executing SQL query: %s", err)
			return models.EnrichedPerson{}, err
		}
		return person, nil
	}
	r.log.Infof("Update person data successfully. %s", utils.StructToString(person))
	return models.EnrichedPerson{}, fmt.Errorf("No rows updated")
}

func (r *Repository) DelPerson(id int64) error {
	r.log.Debug("Repository DelPerson")

	query := "DELETE FROM persons WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		r.log.Errorf("Error executing delete query for person with ID %d: %s", id, err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		r.log.Errorf("No rows deleted for person with ID %d", id)
		return sql.ErrNoRows
	}

	r.log.Infof("Person delete successfully. ID: %d", id)
	return nil
}

func (r *Repository) AddPerson(person models.EnrichedPerson) (models.EnrichedPerson, error) {
	r.log.Debug("Repository AddPerson")

	query := `
		INSERT INTO persons (name, surname, patronymic, age, gender, country_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	r.log.Debugf("Executing SQL query: %s", query)

	err := r.db.Get(&person.ID, query, person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.CountryID)
	if err != nil {
		r.log.Errorf("Error executing SQL query: %s", err)
		return models.EnrichedPerson{}, err
	}

	r.log.Infof("Person added successfully. %s", utils.StructToString(person))

	return person, nil
}

func (r *Repository) GetPersons(filter models.PersonFilter) (models.PersonList, error) {
	r.log.Debug("Repository GetPersons")
	if filter.Limit == 0 {
		filter.Limit = 10
	}
	filterQuery := createFilterQuery(filter)
	sqlQuery := fmt.Sprintf("SELECT * FROM persons %s ORDER BY %s ASC, id ASC LIMIT %d OFFSET %d", filterQuery, *filter.Sort, filter.Limit, filter.Offset)
	r.log.Debugf("Executing SQL query: %s", sqlQuery)

	var persons []models.EnrichedPerson
	err := r.db.Select(&persons, sqlQuery)
	if err != nil {
		r.log.Errorf("Error fetching persons: %s", err)
		return models.PersonList{}, err
	}

	totalCountQuery := fmt.Sprintf("SELECT COUNT(*) FROM persons %s", filterQuery)
	var totalCount int64
	err = r.db.Get(&totalCount, totalCountQuery)
	if err != nil {
		r.log.Errorf("Error fetching total count: %s", err)
		return models.PersonList{}, err
	}

	r.log.Infof("Fetched %d persons successfully", len(persons))

	personList := models.PersonList{
		Persons: persons,
		Offset:  filter.Offset,
		Limit:   filter.Limit,
		Count:   totalCount,
	}

	return personList, nil
}
