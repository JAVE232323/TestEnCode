package db

import (
	"log"
	"test-encode/app"

	"github.com/gocraft/dbr"
)

type PersonRepository struct {
	conn *dbr.Connection
}

func NewPersonRepository(conn *dbr.Connection) *PersonRepository {
	if conn == nil {
		log.Fatal("Ошибка: передана nil-сессия в репозиторий")
	}
	return &PersonRepository{conn: conn}
}

func (r *PersonRepository) GetAll(limit, offset int, search string) ([]app.Person, error) {
	session := r.conn.NewSession(nil)
	var persons []app.Person

	query := session.Select("*").From("persons")

	if search != "" {
		query.Where("first_name ILIKE ? OR last_name ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	_, err := query.Limit(uint64(limit)).Offset(uint64(offset)).Load(&persons)
	return persons, err
}

func (r *PersonRepository) GetById(id int) (*app.Person, error) {
	session := r.conn.NewSession(nil)
	var person app.Person

	err := session.Select("id", "email", "phone", "first_name", "last_name").
		From("persons").
		Where("id = ?", id).
		LoadOne(&person)

	return &person, err
}

func (r *PersonRepository) Create(person *app.Person) error {
	session := r.conn.NewSession(nil)
	_, err := session.InsertInto("persons").Columns("email", "phone", "first_name", "last_name").Record(person).Exec()
	return err
}

func (r *PersonRepository) Update(id int, person *app.Person) error {
	session := r.conn.NewSession(nil)

	_, err := session.Update("persons").
		Set("email", person.Email).
		Set("phone", person.Phone).
		Set("first_name", person.FirstName).
		Set("last_name", person.LastName).
		Where("id = ?", id).
		Exec()
	return err
}

func (r *PersonRepository) Delete(id int) error {
	session := r.conn.NewSession(nil)
	_, err := session.DeleteFrom("persons").Where("id = ?", id).Exec()
	return err
}
