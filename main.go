package main

import (
	"exoplanetservice/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	exoplanets := r.Group("/exoplanets")
	{
		exoplanets.POST("", controllers.AddExoplanet)
		exoplanets.GET("", controllers.ListExoplanets)
		exoplanets.GET("/:id", controllers.GetExoplanetByID)
		exoplanets.PUT("/:id", controllers.UpdateExoplanet)
		exoplanets.DELETE("/:id", controllers.DeleteExoplanet)
	}

	r.GET("/fuel-estimation/:id/:crewCapacity", controllers.FuelEstimation)

	r.Run(":8080")
}
