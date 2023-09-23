package handlers

import (
	"log"
	"net/http"
	

	"github.com/gin-gonic/gin"
	"github.com/micael-ortega/eventos-api/internals"
	"github.com/micael-ortega/eventos-api/models"
)

func GetAllEventos(c *gin.Context) {
	var eventos []models.Evento

	db := internals.OpenDb()

	defer db.Close()

	sqlStmt := "SELECT * FROM evento"

	rows, err := db.Query(sqlStmt)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var evento models.Evento
	for rows.Next() {
		err := rows.Scan(&evento.ID,
			&evento.Data_ini,
			&evento.Data_fim,
			&evento.Duracao,
			&evento.Instrutor.Nome,
			&evento.Curso.Curso)
		if err != nil {
			log.Fatal(err)
		}

		eventos = append(eventos, evento)
	}
	c.JSON(http.StatusOK, eventos)
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
		SELECT e.id, e.data_ini, e.data_fim, e.modalidade, e.duracao,
			   i.nome, AS instrutor_nome, c.id AS curso_id, c.curso AS curso_nome
		FROM evento e
		INNER JOIN intrutor i ON e.instrutor_id = i.id
		INNER JOIN curso c ON e.curso_id = c.id
		WHERE e.id = ?
	`

	row := db.QueryRow(sqlStmt, id)

	var evento models.Evento

	err := row.Scan(
		&evento.ID,
		&evento.Curso.ID,
		&evento.Curso.Curso,
		&evento.Data_ini,
		&evento.Data_fim,
		&evento.Modalidade,
		&evento.Duracao,
		&evento.Instrutor.Nome,
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusOK, evento)
}

func CreateEvento(c *gin.Context) {
	var novoEvento models.Evento

	if err := c.ShouldBindJSON(&novoEvento); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if novoEvento.Data_ini.IsZero() || novoEvento.Data_fim.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Both date fields are required"})
		return
	}

	if novoEvento.Duracao <= 0{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Duration must be greater than zero"})
	}

	db := internals.OpenDb()

	defer db.Close()

	sqlStmt := `
	INSERT INTO evento (
		data_ini, 
		data_fim, 
		modalidade, 
		duracao, 
		instrutor_id, 
		curso_id)
	VALUES (?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(sqlStmt, novoEvento.Data_ini,
		novoEvento.Data_fim, novoEvento.Modalidade,
		novoEvento.Duracao, novoEvento.Instrutor.ID,
		novoEvento.Curso.ID)

	if err != nil {
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusCreated, novoEvento)
}

func UpdateEvento(c *gin.Context) {}

func DeleteEvento(c *gin.Context) {}
