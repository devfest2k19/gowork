package repositories

import (
	"database/sql"
	"github.com/devfest2k19/gowork/session2/9.crud-persistent/pkg/models"
)

type personRepository struct {
	db *sql.DB
}

// NewPersonRepository creates a new repository
func NewPersonRepository(db *sql.DB) *personRepository {
	return &personRepository{
		db,
	}
}

// AddPerson adds a new person to the database
func (pr *personRepository) Add(person models.Person) (int, error) {

	result, err := pr.db.Exec(`
		INSERT INTO phonebook.person
		(first_name, last_name, city, areacode, phone)
		VALUES(?, ?, ?, ?, ?)
	`, person.Firstname, person.Lastname, person.City, person.AreaCode, person.Phone)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// GetAll returns all the person entries that we have in the database
func (pr *personRepository) GetAll() ([]models.Person, error) {

	rows, err := pr.db.Query(`
		SELECT id, first_name, last_name, city, areacode, phone
		FROM phonebook.person
	`)

	if err != nil {
		return nil, err
	}

	persons := []models.Person{}

	for rows.Next() {
		person := models.Person{}
		err = rows.Scan(&person.ID, &person.Firstname, &person.Lastname, &person.City, &person.AreaCode, &person.Phone)
		if err != nil {
			return nil, err
		}

		persons = append(persons, person)

	}

	return persons, nil
}

// GetPersonByID returns a single entry
func (pr *personRepository) GetByID(id int) (bool, models.Person, error) {

	row := pr.db.QueryRow(`
		SELECT id, first_name, last_name, city, areacode, phone
		FROM phonebook.person WHERE id = ?
	`, id)

	person := models.Person{}

	err := row.Scan(&person.ID, &person.Firstname, &person.Lastname, &person.City, &person.AreaCode, &person.Phone)

	if err != nil {

		if err == sql.ErrNoRows {
			return false, models.Person{}, nil
		}

		return false, models.Person{}, err
	}

	return true, person, nil
}

// DeletePerson deletes a person from the database
func (pr *personRepository) Delete(id int) (bool, error) {

	result, err := pr.db.Exec(`
		DELETE FROM phonebook.person
		WHERE id=?;	
	`, id)

	if err != nil {
		return false, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if affectedRows <= 0 {
		return false, nil
	}

	return true, nil
}

// UpdatePerson updates an existing record
func (pr *personRepository) Update(id int, person models.Person) (bool, error) {

	result, err := pr.db.Exec(`
		UPDATE phonebook.person
		SET first_name=?, last_name=?, city=?, areacode=?, phone=?
		WHERE id=?;
		
	`, person.Firstname, person.Lastname, person.City, person.AreaCode, person.Phone, id)

	if err != nil {
		return false, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if affectedRows <= 0 {
		return false, nil
	}

	return true, nil
}
