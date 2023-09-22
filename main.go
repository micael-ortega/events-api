package main

import (
	"./src/handlers/curso_handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/cursos", getCursos)
	router.POST("/cursos", postCursos)
}
