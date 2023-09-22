package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micael-ortega/eventos-api/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/cursos", handlers.GetAllCursos)
	router.POST("/cursos", handlers.CreateCurso)
	router.POST("/deleta/cursos", handlers.DeleteCurso)
	router.Run(":8080")
}
