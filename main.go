package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micael-ortega/events-api/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/course", handlers.GetAllCourses)
	router.POST("/course", handlers.CreateCourse)
	router.DELETE("/course", handlers.DeleteCourse)
	router.GET("/instructor", handlers.GetAllInstructors)
	router.POST("/instructor", handlers.CreateInstructor)
	router.DELETE("/instructor", handlers.DeleteInstructor)
	router.Run(":8080")
}
