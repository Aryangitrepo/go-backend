package main

import (
	"intern/middleware"
	"intern/models"
	"intern/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	ge := gin.Default()
	models.Config()
	nA := ge.Group("/")
	{
		nA.POST("/signup", routes.SignUp)
		nA.POST("/login", routes.Login)
	}
	auth := ge.Group("/auth")
	auth.Use(middleware.Auth)
	{
		auth.GET("/hello", routes.Hello)
		auth.POST("/registerp", routes.RegisterPatient)
		auth.GET("/getallpatients", routes.GetAllPatients)
		auth.GET("/patient", routes.GetPatientByID)
		auth.PATCH("/updatep", routes.UpdatePatient)
		auth.DELETE("/deletep", routes.DeletePatient)
	}
	ge.Run(":8080")
}
