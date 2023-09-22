package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/micael-ortega/eventos-api/internals"
	"github.com/micael-ortega/eventos-api/models"
)

func GetAllCursos(c *gin.Context) {
	var cursos []models.Curso
	db := internals.OpenDb()

	defer db.Close()

	sqlStmt := "SELECT * FROM curso"

	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var curso models.Curso
	for rows.Next() {
		scanErr := rows.Scan(&curso.ID, &curso.Curso)

		if scanErr != nil {
			log.Fatal(scanErr)
			return
		}
		cursos = append(cursos, curso)
	}
	c.IndentedJSON(http.StatusOK, cursos)

}

func CreateCurso(c *gin.Context) {
	var novoCurso models.Curso

	if err := c.ShouldBindJSON(&novoCurso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := internals.OpenDb()

	defer db.Close()

	sqlStmt := "INSERT INTO curso (curso) VALUES (?)"

	_, err := db.Exec(sqlStmt, novoCurso.Curso)

	if err != nil {
		log.Fatal(err)
		return
	}

	c.IndentedJSON(http.StatusCreated, novoCurso)
}

func DeleteCurso(c *gin.Context) {
	var id int16 

	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := internals.OpenDb()

	defer db.Close()

	sqlStmt := "DELETE FROM curso WHERE id = (?)"

	_, err := db.Exec(sqlStmt, id)

	if err != nil {
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusNoContent, "Curso deletado")
}
