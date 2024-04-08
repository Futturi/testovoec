package service

import (
	"github.com/Futturi/testovoe/internal/models"
	"github.com/Futturi/testovoe/internal/repository"
)

type Service struct {
	CarsService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{CarsService: NewCarService(repo.CarsRepo)}
}

type CarsService interface {
	GetNomers(page int, query map[string]string) ([]models.Car, error)
	Delete(nomer models.Car) error
	Put(nomer models.CarUpdate) error
	Insert(nomers []models.Car) error
}
