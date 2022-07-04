package services

import (
	"go-ddd/domain/customer"
	"go-ddd/domain/customer/memory"

	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
}

func NewOrderService(confs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	for _, conf := range confs {
		err := conf(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) error {
	_, err := o.customers.Get(customerID)
	if err != nil {
		return err
	}
	return nil
}
