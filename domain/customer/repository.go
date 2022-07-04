package customer

import (
	"errors"
	"go-ddd/aggregate"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("customer not found")
	ErrFailedToAddCustomer = errors.New("Failed to add customer")
	ErrUpdateCustomer      = errors.New("Failed to update customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer,error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}