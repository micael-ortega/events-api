package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/micael-ortega/eventos-api/models"
	"log"
	"net/http"
)

func CreateCurso(c *gin.Context) {

	var novoCurso models.Curso

	db, err := sql.Open("sqlite3", "../../database.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sqlStmt := "INSERT INTO TABLE curso(curso) VALUES(?)"

	_, err = db.Exec(sqlStmt, novoCurso.Curso)

	c.IndentedJSON(http.StatusCreated, novoCurso)
}

func GetAllCursos(c *gin.Context) {
	var result []models.Curso
	db, err := sql.Open("sqlite3", "../../database.db")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	sqlStmt := "SELECT * FROM curso"

	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {

	}

}
