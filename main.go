package main

// TODO: ADD MORE ENDPOINTS (CREATE, UPDATE, DELETE)
// TODO SEPARATE HANDLERS AND JSON STRUCTS INTO DIFFERENT FILES
// TODO: ADD DATABASE INTEGRATION
// TODO: ADD TESTS

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	ID    int    `json:"id"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  string `json:"year"`
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.GET("/cars/:id", getCar)
	r.POST("/cars", createCar)
	r.PUT("/cars/:id", updateCar)
	r.DELETE("/cars/:id", deleteCar)
	r.Run(":8080")
}

func getCar(c *gin.Context) {
	id := c.Param("id")
	_ = id // In a real application, you would use this ID to fetch the car from a database
	car := Car{
		ID:    1,
		Make:  "Toyota",
		Model: "Camry",
		Year:  "2025",
	}
	c.JSON(http.StatusOK, car)
}

func createCar(c *gin.Context) {
	var newCar Car
	if err := c.ShouldBindJSON(&newCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// In a real application, you would save the new car to a database here
	newCar.ID = 2 // This is just a placeholder
	c.JSON(http.StatusCreated, newCar)
}

func updateCar(c *gin.Context) {
	id := c.Param("id")
	var updatedCar Car
	if err := c.ShouldBindJSON(&updatedCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// In a real application, you would update the car in the database here
	c.JSON(http.StatusOK, gin.H{"message": "Car with ID " + id + " updated", "car": updatedCar})
}

func deleteCar(c *gin.Context) {
	id := c.Param("id")
	// In a real application, you would delete the car from the database here
	c.JSON(http.StatusOK, gin.H{"message": "Car with ID " + id + " deleted"})
}
