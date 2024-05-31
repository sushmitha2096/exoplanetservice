package controllers

import (
	"exoplanetservice/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FuelEstimation(c *gin.Context) {
	id := c.Param("id")
	crewCapacityStr := c.Param("crewCapacity")

	crewCapacity, err := strconv.Atoi(crewCapacityStr)
	if err != nil || crewCapacity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid crew capacity"})
		return
	}

	exoplanet, exists := models.Exoplanets[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exoplanet not found"})
		return
	}

	distance := exoplanet.Distance
	radius := exoplanet.Radius
	var gravity float64

	switch exoplanet.Type {
	case "GasGiant":
		gravity = 0.5 / (radius * radius)
	case "Terrestrial":
		if exoplanet.Mass == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Terrestrial planet missing mass information"})
			return
		}
		gravity = *exoplanet.Mass / (radius * radius)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown exoplanet type"})
		return
	}

	fuel := float64(distance) / (gravity * gravity) * float64(crewCapacity)
	c.JSON(http.StatusOK, gin.H{"fuel_estimation": fuel})
}
