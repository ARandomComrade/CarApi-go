package repositories

import (
	"example/car-api/models"
)

type CarRepository struct{}

func (r *CarRepository) GetCar(id string) (models.Car, error) {
	_ = id // In a real application, you would use this ID to fetch the car from a database
	return models.Car{
		ID:    1,
		Make:  "Toyota",
		Model: "Camry",
		Year:  "2025",
	}, nil

}

func (r *CarRepository) CreateCar(newCar models.Car) (models.Car, error) {
	_ = newCar    // In a real application, you would save the new car to a database here
	newCar.ID = 2 // This is just a placeholder
	return newCar, nil
}

func (r *CarRepository) UpdateCar(id string, updatedCar models.Car) (models.Car, error) {
	_ = id // In a real application, you would use this ID to update the car in the database
	return updatedCar, nil
}

func (r *CarRepository) DeleteCar(id string) error {
	_ = id // In a real application, you would use this ID to delete the car from the database
	return nil
}
