package main

import (
	"intern/middleware"
	"intern/models"
	"intern/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func setupServer() *gin.Engine {
	// Gin Mode
	gin.SetMode(gin.DebugMode)
	if err := models.Config(); err != nil {
		log.Fatal(err)
	}
	// Creates the Gin Engine
	engine := gin.New()

	// Setup the API Routes

	setupUserRoutes(engine)

	// Return engine
	return engine
}

func setupUserRoutes(ge *gin.Engine) {
	nA := ge.Group("/")
	{
		services := models.User{}
		nA.POST("/signup", routes.SignUp(&services))
		nA.POST("/login", routes.Login(&services))
	}

	auth := ge.Group("/auth")
	auth.Use(middleware.Auth)
	{
		service := models.PatientData{}
		auth.GET("/hello", routes.Hello)
		auth.POST("/registerp", routes.RegisterPatient(&service))
		auth.GET("/getallpatients", routes.GetAllPatients(&service))
		auth.GET("/patient", routes.GetPatientByID(&service))
		auth.PATCH("/updatep", routes.UpdatePatient(&service))
		auth.DELETE("/deletep", routes.DeletePatient(&service))
	}
}

func main() {
	setupServer().Run(":8080")
}
