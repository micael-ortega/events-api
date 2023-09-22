package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micael-ortega/eventos-api/internals"
	"github.com/micael-ortega/eventos-api/models"
)

func GetAllInstrutores(c *gin.Context){
	var instrutores []models.Instrutor

	db := internals.OpenDb()

	defer db.Close()

	sqlStmt := "SELECT * FROM instrutor"

	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	
	var instrutor models.Instrutor
	for rows.Next() {
		err := rows.Scan(&instrutor.ID, &instrutor.Nome)
		
		if err != nil {
			log.Fatal(err)
		}
		instrutores = append(instrutores, instrutor)
	}
	c.JSON(http.StatusOK, instrutores)
}