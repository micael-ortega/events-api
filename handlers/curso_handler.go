package main

import (
	"database/sql"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func createCurso(c *gin.Context) {

	var novoCurso curso

	db, err := sql.Open("sqlite3", "../../database.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sqlStmt := "INSERT INTO TABLE curso(curso) VALUES(?)"

	_, err = db.Exec(sqlStmt, novoCurso.curso)
	

	c.IndentedJSON(http.StatusCreated)
}

func getAllCursos(c *gin.Context){
	var result []curso
	db, err := sql.Open("sqlite3", "../../database.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sqlStmt := "SELECT * FROM curso"

	_, err = db.Query(sqlStmt)

	if err != nil {
		log.Fatal(err)
	}
}
