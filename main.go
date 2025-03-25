package main

import (
	"os"
	"test-encode/internal/db"
	"test-encode/internal/handlers"
	"test-encode/internal/logic"

	"github.com/gocraft/dbr"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {

	logger := logrus.New()
	logger.Out = os.Stdout

	dbConn, err := dbr.Open("postgres", "user=postgres password=123 dbname=persondb sslmode=disable", nil)
	if err != nil {
		logger.Fatal("Ощибка подключения к бд", err)
	}
	defer dbConn.Close()

	dbRepo := db.NewPersonRepository(dbConn)
	PersonLogic := logic.NewPersonLogic(dbRepo, logger)
	PersonHandler := handlers.NewPersonHandler(PersonLogic)

	e := echo.New()

	e.GET("/person", PersonHandler.GetPersons)
	e.GET("/person/:id", PersonHandler.GetPerson)
	e.POST("/person", PersonHandler.CreatePerson)
	e.PUT("/person/:id", PersonHandler.UpdatePerson)
	e.DELETE("/person/:id", PersonHandler.DeletePerson)

	if err := e.Start(":8080"); err != nil {
		logger.Fatal(err)
	}
}
