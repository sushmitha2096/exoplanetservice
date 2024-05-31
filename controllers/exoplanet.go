package controllers

import (
	"exoplanetservice/models"
	"exoplanetservice/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddExoplanet(c *gin.Context) {
	var exoplanet models.Exoplanet
	if err := c.ShouldBindJSON(&exoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validation
	if err := utils.ValidateExoplanet(exoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exoplanet.ID = uuid.New().String()
	models.Exoplanets[exoplanet.ID] = exoplanet
	c.JSON(http.StatusCreated, exoplanet)
}

func ListExoplanets(c *gin.Context) {
	exoplanets := make([]models.Exoplanet, 0, len(models.Exoplanets))
	for _, exoplanet := range models.Exoplanets {
		exoplanets = append(exoplanets, exoplanet)
	}
	c.JSON(http.StatusOK, exoplanets)
}

func GetExoplanetByID(c *gin.Context) {
	id := c.Param("id")
	exoplanet, exists := models.Exoplanets[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exoplanet not found"})
		return
	}
	c.JSON(http.StatusOK, exoplanet)
}

func UpdateExoplanet(c *gin.Context) {
	id := c.Param("id")
	var updatedExoplanet models.Exoplanet
	if err := c.ShouldBindJSON(&updatedExoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, exists := models.Exoplanets[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exoplanet not found"})
		return
	}

	// Validation
	if err := utils.ValidateExoplanet(updatedExoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedExoplanet.ID = id
	models.Exoplanets[id] = updatedExoplanet
	c.JSON(http.StatusOK, updatedExoplanet)
}

func DeleteExoplanet(c *gin.Context) {
	id := c.Param("id")
	if _, exists := models.Exoplanets[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exoplanet not found"})
		return
	}
	delete(models.Exoplanets, id)
	c.Status(http.StatusNoContent)
}
