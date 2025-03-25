package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"test-encode/app"
	"test-encode/internal/logic"

	"github.com/labstack/echo/v4"
)

type PersonHandler struct {
	logic *logic.PersonLogic
}

func NewPersonHandler(logic *logic.PersonLogic) *PersonHandler {
	return &PersonHandler{logic: logic}
}

func (h *PersonHandler) GetPersons(c echo.Context) error {

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	search := c.QueryParam("search")

	if limit == 0 {
		limit = 10
	}

	persons, err := h.logic.GetAll(limit, offset, search)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, persons)
}

func (h *PersonHandler) GetPerson(c echo.Context) error {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	person, err := h.logic.GetById(id)
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
	if err := h.logic.Create(&person); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) UpdatePerson(c echo.Context) error {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	fmt.Print(id)
	var person app.Person
	if err := c.Bind(&person); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	person.Id = id
	if err := h.logic.Update(id, &person); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) DeletePerson(c echo.Context) error {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	if err := h.logic.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusNoContent, nil)
}
