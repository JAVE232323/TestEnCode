package logic

import (
	"test-encode/app"
	"test-encode/internal/db"

	"github.com/sirupsen/logrus"
)

type PersonLogic struct {
	repo   *db.PersonRepository
	logger *logrus.Logger
}

func NewPersonLogic(repo *db.PersonRepository, logger *logrus.Logger) *PersonLogic {
	return &PersonLogic{repo: repo, logger: logger}
}

func (pl *PersonLogic) GetAll(limit, offset int, search string) ([]app.Person, error) {
	pl.logger.Infof("Выборка лиц: лимит=%d, offset=%d, поиск=%s", limit, offset, search)
	return pl.repo.GetAll(limit, offset, search)
}

func (pl *PersonLogic) GetById(id int) (*app.Person, error) {
	pl.logger.Infof("Выборка по id=%d", id)
	return pl.repo.GetById(id)
}

func (pl *PersonLogic) Create(person *app.Person) error {
	pl.logger.Infof("Добавление человека: %v", person)
	return pl.repo.Create(person)
}

func (pl *PersonLogic) Update(id int, person *app.Person) error {
	pl.logger.Infof("Обновлене человека id=%d", id)
	return pl.repo.Update(id, person)
}

func (pl *PersonLogic) Delete(id int) error {
	pl.logger.Infof("Удаленеи человека: id=%d", id)
	return pl.repo.Delete(id)
}
