package main

// TODO: ADD TESTS
// TODO: ADD DATABASE INTEGRATION

import (
	"example/car-api/handlers"
	"example/car-api/repositories"
	"example/car-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Initialize repositories, services, and handlers
	repo := &repositories.CarRepository{}
	service := &services.CarService{Repo: repo}
	carHandler := &handlers.CarHandler{Service: service}

	r.GET("/cars/:id", carHandler.GetCar)
	r.POST("/cars", carHandler.CreateCar)
	r.PUT("/cars/:id", carHandler.UpdateCar)
	r.DELETE("/cars/:id", carHandler.DeleteCar)
	r.Run(":8080")
}
