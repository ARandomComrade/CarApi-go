package handlers

import (
	"example/car-api/models"
	"example/car-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CarHandler struct {
	Service *services.CarService
}

func (h *CarHandler) GetCar(c *gin.Context) {
	id := c.Param("id")
	car, _ := h.Service.GetCar(id) // This is just a placeholder
	c.JSON(http.StatusOK, car)
}

func (h *CarHandler) CreateCar(c *gin.Context) {
	var newCar models.Car
	if err := c.ShouldBindJSON(&newCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	// In a real application, you would create the car in the database here
	h.Service.CreateCar(newCar) // This is just a placeholder
	c.JSON(http.StatusCreated, gin.H{"message": "Car created", "car": newCar})
}

func (h *CarHandler) UpdateCar(c *gin.Context) {
	id := c.Param("id")
	var updatedCar models.Car
	if err := c.ShouldBindJSON(&updatedCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// In a real application, you would update the car in the database here
	h.Service.UpdateCar(id, updatedCar) // This is just a placeholder
	c.JSON(http.StatusOK, gin.H{"message": "Car with ID " + id + " updated", "car": updatedCar})
}

func (h *CarHandler) DeleteCar(c *gin.Context) {
	id := c.Param("id")
	// In a real application, you would delete the car from the database here
	h.Service.DeleteCar(id) // This is just a placeholder
	c.JSON(http.StatusOK, gin.H{"message": "Car with ID " + id + " deleted"})
}
