package main

import (
	"database/sql"
	"log"
	"test-encode/db"
	"test-encode/handlers"

	"github.com/gocraft/dbr"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	dbConn, err := sql.Open("postgres", "user=postgres password=123 dbname=persondb sslmode=disable")
	if err != nil {
		log.Fatal("Ощибка подключения к бд", err)
	}

	session := &dbr.Session{Connection: &dbr.Connection{DB: dbConn}, EventReceiver: &dbr.NullEventReceiver{}}

	dbRepo := db.NewPersonRepository(session)
	PersonHandler := handlers.NewPersonHandler(dbRepo)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/person", PersonHandler.GetPersons)
	e.GET("/person/:id", PersonHandler.GetPerson)
	e.POST("/person", PersonHandler.CreatePerson)
	e.PUT("/person/:ig", PersonHandler.UpdatePerson)
	e.DELETE("/person/:id", PersonHandler.DeletePerson)

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
