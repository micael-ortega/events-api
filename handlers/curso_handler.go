package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/micael-ortega/eventos-api/models"
)

func CreateCurso(c *gin.Context) {

	var novoCurso models.Curso

	db, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sqlStmt := "INSERT INTO TABLE curso(curso) VALUES(?)"

	_, err = db.Exec(sqlStmt, novoCurso.Curso)

	if err != nil {
		log.Fatal(err)
		return
	}

	c.IndentedJSON(http.StatusCreated, novoCurso)
}

func GetAllCursos(c *gin.Context) {
	var cursos []models.Curso
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	sqlStmt := "SELECT * FROM curso"

	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Fatal(err)
		return
	}

	for rows.Next() {
		var curso models.Curso
		scanErr := rows.Scan(&curso)

		if scanErr != nil {
			log.Fatal(scanErr)
			return
		}
		cursos = append(cursos, curso)
	}
	c.JSON(http.StatusOK, cursos)

}
