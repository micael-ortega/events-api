package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/micael-ortega/events-api/internals"
	"github.com/micael-ortega/events-api/models"
)

func GetAllEvents(c *gin.Context) {
	var events []models.EventResponse

	db := internals.OpenDb()

	defer db.Close()

	sqlStmt := `
		SELECT e.id, 
			e.begin_date, 
			e.end_date, 
			e.modality, 
			e.duration,
			i.id AS instructor_id, 
			i.name AS instructor_name, 
			c.id AS course_id, 
			c.course AS course_name
		FROM event e
		JOIN instructor i ON e.instructor_id = i.id
		JOIN course c ON e.course_id = c.id
	`
	rows, err := db.Query(sqlStmt)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	defer db.Close()


	var event models.EventResponse
	for rows.Next() {
		err := rows.Scan(
			&event.ID,
			&event.Begin_date,
			&event.End_date,
			&event.Modality,
			&event.Duration,
			&event.Instructor.ID,
			&event.Instructor.Name,
			&event.Course.ID,
			&event.Course.Course,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		events = append(events, event)
	}
	c.JSON(http.StatusOK, events)
}

func GetEventById(c *gin.Context) {
	eventID := c.Param("id")

	id, err := strconv.Atoi(eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	db := internals.OpenDb()
	defer db.Close()

	sqlStmt := `
		SELECT e.id, e.begin_date, e.end_date, e.modality, e.duration,
			   i.id AS instructor_id, 
			   i.name AS instructor_name, 
			   c.id AS course_id, 
			   c.course AS course_name
		FROM event e
		INNER JOIN instructor i ON e.instructor_id = i.id
		INNER JOIN course c ON e.course_id = c.id
		WHERE e.id = ?
	`

	row := db.QueryRow(sqlStmt, id)

	var event models.EventResponse

	err = row.Scan(
		&event.ID,
		&event.Begin_date,
		&event.End_date,
		&event.Modality,
		&event.Duration,
		&event.Instructor.ID,
		&event.Instructor.Name,
		&event.Course.ID,
		&event.Course.Course,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, event)
}

func CreateEvent(c *gin.Context) {
	var newEvent models.Event

	if err := c.ShouldBindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newEvent.Begin_date == "" || newEvent.End_date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both date fields are required"})
		return
	}

	if newEvent.Duration <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Duration must be greater than zero"})
		return
	}

	if newEvent.Instructor == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Instructor must be assigned to event"})
		return
	}

	if newEvent.Course == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course must be assigned to event"})
		return
	}

	db := internals.OpenDb()

	defer db.Close()

	instructorExists := CheckIfInstructorExists(newEvent.Instructor)

	if !instructorExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Instructor with specified ID does not exist"})
		return
	}

	courseExists := CheckIfCourseExists(newEvent.Course)

	if !courseExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course with specified ID does not exist"})
		return
	}

	sqlStmt := `
		INSERT INTO event (
			begin_date, 
			end_date, 
			modality, 
			duration, 
			instructor_id, 
			course_id)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(sqlStmt, newEvent.Begin_date,
		newEvent.End_date, newEvent.Modality,
		newEvent.Duration, newEvent.Instructor,
		newEvent.Course)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newEvent)
}

func UpdateEvent(c *gin.Context) {
	// TODO 
}

func DeleteEvent(c *gin.Context) {
	var request struct {
		ID int `json:"id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := internals.OpenDb()

	defer db.Close()

	sqlStmt := "DELETE FROM event WHERE id = (?)"

	_, err := db.Exec(sqlStmt, request.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
