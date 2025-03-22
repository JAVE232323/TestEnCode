package db

import (
	"errors"
	"log"
	"test-encode/app"

	"github.com/gocraft/dbr"
)

type PersonRepository struct {
	session *dbr.Session
}

func NewPersonRepository(session *dbr.Session) *PersonRepository {
	if session == nil {
		log.Fatal("Ошибка: передана nil-сессия в репозиторий")
	}
	return &PersonRepository{session: session}
}

func (r *PersonRepository) GetAll() ([]app.Person, error) {
	var persons []app.Person
	session := r.session.Select("id", "email", "phone", "first_name", "last_name").
		From("persons")
	_, err := session.Load(&persons)
	return persons, err
}

func (r *PersonRepository) GetById(id int) (*app.Person, error) {
	var person app.Person
	session := r.session.Select("id", "email", "phone", "first_name", "last_name").
		From("persons").
		Where("id = ?", id)
	_, err := session.Load(&person)
	return &person, err
}

func (r *PersonRepository) Create(person *app.Person) error {
	if person == nil {
		return errors.New("пустая структура Person")
	}
	return r.session.InsertInto("persons").Columns("email", "phone", "first_name", "last_name").Record(person).Returning("id").Load(&person.Id)
}

func (r *PersonRepository) Update(id int, person *app.Person) error {
	_, err := r.session.Update("persons").
		Set("email", person.Email).
		Set("phone", person.Phone).
		Set("first_name", person.FirstName).
		Set("last_name", person.LastName).
		Where("id = ?", id).
		Exec()
	return err
}

func (r *PersonRepository) Delete(id int) error {
	_, err := r.session.DeleteFrom("persons").Where("id = ?", id).Exec()
	return err
}
