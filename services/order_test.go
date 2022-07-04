package services

import (
	"go-ddd/aggregate"
	"testing"

	"github.com/google/uuid"
)

func init_products(t *testing.T) []aggregate.Product {
	wine, err := aggregate.NewProduct("Wine", "Red Wine", 10.0)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := aggregate.NewProduct("Peenuts", "Peenuts", 10.0)
	if err != nil {
		t.Error(err)
	}
	beer, err := aggregate.NewProduct("Beer", "Beer", 10.0)
	if err != nil {
		t.Error(err)
	}
	products := []aggregate.Product{wine, peenuts, beer}
	return products
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)
	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Error(err)
	}
	cust, err := aggregate.NewCustomer("Jhon")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	_, err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}
