package handlers

import (
	"log"
	"net/http"
	

	"github.com/gin-gonic/gin"
	"github.com/micael-ortega/events-api/internals"
	"github.com/micael-ortega/events-api/models"
)

func GetAllEventos(c *gin.Context) {
	var events []models.Event

	db := internals.OpenDb()

	defer db.Close()

	sqlStmt := "SELECT * FROM event"

	rows, err := db.Query(sqlStmt)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var event models.Event
	for rows.Next() {
		err := rows.Scan(&event.ID,
			&event.Begin_date,
			&event.End_date,
			&event.Duration,
			&event.Instructor.Name,
			&event.Course.Course)
		if err != nil {
			log.Fatal(err)
		}

		events = append(events, event)
	}
	c.JSON(http.StatusOK, events)
}

func GetEventoById(c *gin.Context) {
	var id int16

	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := internals.OpenDb()

	defer db.Close()

	sqlStmt := `
		SELECT e.id, e.begin_date, e.end_date, e.modality, e.duration,
			   i.name, AS instructor_name, c.id AS curso_id, c.course AS course_name
		FROM event e
		INNER JOIN instructor i ON e.instructor_id = i.id
		INNER JOIN course c ON e.course_id = c.id
		WHERE e.id = ?
	`

	row := db.QueryRow(sqlStmt, id)

	var event models.Event

	err := row.Scan(
		&event.ID,
		&event.Course.ID,
		&event.Course.Course,
		&event.Begin_date,
		&event.End_date,
		&event.Modality,
		&event.Duration,
		&event.Instructor.Name,
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusOK, event)
}

func CreateEvento(c *gin.Context) {
	var newEvent models.Event

	if err := c.ShouldBindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if newEvent.Begin_date.IsZero() || newEvent.End_date.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Both date fields are required"})
		return
	}

	if newEvent.Duration <= 0{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Duration must be greater than zero"})
	}

	db := internals.OpenDb()

	defer db.Close()

	sqlStmt := `
	INSERT INTO event (
		begin_date, 
		end_date, 
		modality, 
		duration, 
		instrutor_id, 
		curso_id)
	VALUES (?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(sqlStmt, newEvent.Begin_date,
		newEvent.End_date, newEvent.Modality,
		newEvent.Duration, newEvent.Instructor.ID,
		newEvent.Course.ID)

	if err != nil {
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusCreated, newEvent)
}

func UpdateEvento(c *gin.Context) {}

func DeleteEvento(c *gin.Context) {}
