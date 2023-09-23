package handlers

import (
	"net/http"
	"strconv"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	defer db.Close()

	var instructor models.Instructor
	for rows.Next() {
		err := rows.Scan(&instructor.ID, &instructor.Name)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		instructos = append(instructos, instructor)
	}
	c.JSON(http.StatusOK, instructos)
}

func GetInstructorById(c *gin.Context) {
	instructorID := c.Param("id")

	id, err := strconv.Atoi(instructorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	db := internals.OpenDb()
	defer db.Close()

	sqlStmt := `
		SELECT * FROM instructor WHERE id = ?
	`

	row := db.QueryRow(sqlStmt, id)

	var instructor models.Instructor

	err = row.Scan(
		&instructor.ID,
		&instructor.Name,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, instructor)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func CheckIfInstructorExists(id int)(bool){
	var instructorExists bool
	db := internals.OpenDb()

	defer db.Close()
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM instructor WHERE id = ?)", id).Scan(&instructorExists)
	if err != nil {
		return false
	}
	
	if !instructorExists {
		return false
	}
	return true
}