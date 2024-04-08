package service

import (
	"github.com/Futturi/testovoe/internal/models"
	"github.com/Futturi/testovoe/internal/repository"
)

type Car_Service struct {
	repo repository.CarsRepo
}

func NewCarService(repo repository.CarsRepo) *Car_Service {
	return &Car_Service{repo: repo}
}

func (a *Car_Service) GetNomers(page int, query map[string]string) ([]models.Car, error) {
	return a.repo.GetNomers(page, query)
}

func (a *Car_Service) Delete(nomer models.Car) error {
	return a.repo.Delete(nomer)
}
func (a *Car_Service) Put(nomer models.CarUpdate) error {
	return a.repo.Put(nomer)
}
func (a *Car_Service) Insert(nomers []models.Car) error {
	return a.repo.Insert(nomers)
}
