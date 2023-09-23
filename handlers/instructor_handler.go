package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micael-ortega/events-api/internals"
	"github.com/micael-ortega/events-api/models"
)

func GetAllInstructors(c *gin.Context) {
	var instructos []models.Instructor

	db := internals.OpenDb()

	defer db.Close()

	sqlStmt := "SELECT * FROM instructor"

	rows, err := db.Query(sqlStmt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	}

	defer db.Close()

	var instructor models.Instructor
	for rows.Next() {
		err := rows.Scan(&instructor.ID, &instructor.Name)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		}
		instructos = append(instructos, instructor)
	}
	c.JSON(http.StatusOK, instructos)
}

func CreateInstructor(c *gin.Context) {
	var newInstructor models.Instructor

	if err := c.ShouldBindJSON(&newInstructor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := internals.OpenDb()

	defer db.Close()

	sqlStmt := "INSERT INTO instructor (name) VALUES (?)"

	_, err := db.Exec(sqlStmt, newInstructor.Name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newInstructor)

}

func DeleteInstructor(c *gin.Context) {
	var request struct {
		ID int `json:"id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := internals.OpenDb()

	defer db.Close()

	sqlStmt := "DELETE FROM instructor WHERE id = (?)"

	_, err := db.Exec(sqlStmt, request.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
