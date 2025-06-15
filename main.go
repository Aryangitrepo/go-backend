package main

import (
	"intern/middleware"
	"intern/models"
	"intern/routes"

	"github.com/gin-gonic/gin"
)

/*func setupServer() *gin.Engine {
	// Gin Mode
	gin.SetMode(gin.ReleaseMode)

	// Creates the Gin Engine
	engine := gin.New()

	// Setup the API Routes

	setupUserRoutes(engine)
	auth := engine.Group("/auth")
	auth.Use(middleware.Auth)
	{
		auth.GET("/hello", routes.Hello)
		auth.POST("/registerp", routes.RegisterPatient)
		auth.GET("/getallpatients", routes.GetAllPatients)
		auth.GET("/patient", routes.GetPatientByID)
		auth.PATCH("/updatep", routes.UpdatePatient)
		auth.DELETE("/deletep", routes.DeletePatient)
	}s

	// Return engine
	return engine
}*/

func setupUserRoutes(ge *gin.Engine) {
	nA := ge.Group("/")
	{
		services := models.Userfunc{}
		nA.POST("/signup", routes.SignUp(&services))
		nA.POST("/login", routes.Login(&services))
	}
}

func main() {
	ge := gin.Default()
	models.Config()
	setupUserRoutes(ge)
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
	//setupserver().Run(":8080")
}
