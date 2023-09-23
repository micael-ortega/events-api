package handlers

import (
	"log"
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
		log.Fatal(err)
	}

	defer rows.Close()

	var course models.Course
	for rows.Next() {
		scanErr := rows.Scan(&course.ID, &course.Course)

		if scanErr != nil {
			log.Fatal(scanErr)
			return
		}
		courses = append(courses, course)
	}
	c.IndentedJSON(http.StatusOK, courses)

}

func GetCourseById(c *gin.Context){
	var request struct {
		ID uint16 `json:"id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
	}
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

	var novoCourse models.Course

	novoCourse.Course = strings.TrimSpace(request.Course)

	_, err := db.Exec(sqlStmt, strings.ToTitle(novoCourse.Course))

	if err != nil {
		log.Fatal(err)
		return
	}

	c.IndentedJSON(http.StatusCreated, novoCourse)
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

	sqlStmt := "DELETE FROM course WHERE id = (?)"

	_, err := db.Exec(sqlStmt, request.ID)

	if err != nil {
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

