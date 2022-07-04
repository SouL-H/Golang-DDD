package product

import (
	"errors"
	"go-ddd/aggregate"

	"github.com/google/uuid"
)


var (
	ErrProductNotFound    = errors.New("product not found")
	ErrProductAlreadyExist = errors.New("product already exist")
)

type ProductReporitory interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}