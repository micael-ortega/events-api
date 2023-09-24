package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micael-ortega/events-api/handlers"
	"github.com/micael-ortega/events-api/internals"
)

func main() {
	router := gin.Default()
	router.Use(internals.CORSMiddleware())
	router.GET("/course", handlers.GetAllCourses)
	router.POST("/course", handlers.CreateCourse)
	router.DELETE("/course", handlers.DeleteCourse)
	router.GET("/instructor", handlers.GetAllInstructors)
	router.POST("/instructor", handlers.CreateInstructor)
	router.DELETE("/instructor", handlers.DeleteInstructor)
	router.GET("/event", handlers.GetAllEvents)
	router.POST("/event", handlers.CreateEvent)
	router.GET("/event/:id", handlers.GetEventById)
	router.DELETE("/event", handlers.DeleteEvent)
	router.Run(":8080")
}
