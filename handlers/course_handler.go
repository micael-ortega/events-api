package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/micael-ortega/events-api/internals"
	"github.com/micael-ortega/events-api/models"
)

func GetAllCourses(c *gin.Context) {
	var courses []models.Course
	db := internals.OpenDb()

	defer db.Close()

	sqlStmt := "SELECT * FROM course"

	rows, err := db.Query(sqlStmt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	defer rows.Close()

	var course models.Course
	for rows.Next() {
		err := rows.Scan(&course.ID, &course.Course)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		courses = append(courses, course)
	}
	c.IndentedJSON(http.StatusOK, courses)
}

func GetCourseById(c *gin.Context) {
	var request struct {
		ID int `json:"id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
	}

	db := internals.OpenDb()

	defer db.Close()

	sqlStmt := "SELECT * FROM course WHERE id = ?"

	row := db.QueryRow(sqlStmt, request.ID)

	var course models.Course

	err := row.Scan(
		&course.ID,
		&course.Course,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, course)
}

func CreateCourse(c *gin.Context) {
	var request struct {
		Course string `json:"course"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Course == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course is required"})
		return
	}

	db := internals.OpenDb()

	defer db.Close()

	sqlStmt := "INSERT INTO course (course) VALUES (?)"

	var newCourse models.Course

	newCourse.Course = strings.TrimSpace(request.Course)

	_, err := db.Exec(sqlStmt, strings.ToTitle(newCourse.Course))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newCourse)
}

func DeleteCourse(c *gin.Context) {
	var request struct {
		ID int `json:"id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := internals.OpenDb()

	defer db.Close()

	courseExists := CheckIfCourseExists(request.ID)

	if !courseExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course with specified ID does not exist"})
		return
	}

	sqlStmt := "DELETE FROM course WHERE id = (?)"

	_, err := db.Exec(sqlStmt, request.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func CheckIfCourseExists(id int)(bool){
	var courseExists bool
	db := internals.OpenDb()

	defer db.Close()
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM course WHERE id = ?)", id).Scan(&courseExists)
	if err != nil {
		return false
	}
	
	if !courseExists {
		return false
	}
	return true
}
