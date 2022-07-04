package services

import (
	"context"
	"go-ddd/aggregate"
	"go-ddd/domain/customer"
	cust "go-ddd/domain/customer/memory"
	"go-ddd/domain/customer/mongo"
	"go-ddd/domain/product"
	prod "go-ddd/domain/product/memory"
	"log"

	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers     customer.CustomerRepository
	products      product.ProductReporitory
	mongoCustomer mongo.MongoRepository
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
	cr := cust.New()
	return WithCustomerRepository(cr)
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prod.New()
		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}
		os.products = pr
		return nil

	}

}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}
	var products []aggregate.Product
	var price float64
	for _, id := range productIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		price += p.GetPrice()
	}
	log.Printf("Customer %s ordered %d products for %f", c.GetName(), len(products), price)
	return price, nil
}

func WithMongoCustomerRepository(connectionString string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.New(context.Background(), connectionString)
		if err != nil {
			return err
		}
		os.mongoCustomer = *cr
		return nil
	}
}
