package handlers

import (
	"net/http"
	"strconv"
	"test-encode/app"
	"test-encode/db"

	"github.com/labstack/echo/v4"
)

type PersonHandler struct {
	repo *db.PersonRepository
}

func NewPersonHandler(repo *db.PersonRepository) *PersonHandler {
	return &PersonHandler{repo: repo}
}

func (h *PersonHandler) GetPersons(c echo.Context) error {
	persons, err := h.repo.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, persons)
}

func (h *PersonHandler) GetPerson(c echo.Context) error {
	id, _ := strconv.Atoi("id")
	person, err := h.repo.GetById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) CreatePerson(c echo.Context) error {
	var person app.Person
	if err := c.Bind(&person); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := h.repo.Create(&person); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) UpdatePerson(c echo.Context) error {
	id, _ := strconv.Atoi("id")
	var person app.Person
	if err := c.Bind(&person); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	person.Id = id
	if err := h.repo.Update(&person); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) DeletePerson(c echo.Context) error {
	id, _ := strconv.Atoi("id")
	if err := h.repo.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
