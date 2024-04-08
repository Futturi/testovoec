package repository

import (
	"github.com/Futturi/testovoe/internal/models"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	CarsRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{CarsRepo: NewCars_Repo(db)}
}

type CarsRepo interface {
	GetNomers(page int, query map[string]string) ([]models.Car, error)
	Delete(nomer models.Car) error
	Put(nomer models.CarUpdate) error
	Insert(nomers []models.Car) error
}
