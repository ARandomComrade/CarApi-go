package services

import (
	"example/car-api/models"
	"example/car-api/repositories"
)

type CarService struct {
	Repo *repositories.CarRepository
}

func (s *CarService) GetCar(id string) (models.Car, error) {
	return s.Repo.GetCar(id)
}

func (s *CarService) CreateCar(newCar models.Car) (models.Car, error) {
	return s.Repo.CreateCar(newCar)
}

func (s *CarService) UpdateCar(id string, updatedCar models.Car) (models.Car, error) {
	return s.Repo.UpdateCar(id, updatedCar)
}

func (s *CarService) DeleteCar(id string) error {
	return s.Repo.DeleteCar(id)
}
