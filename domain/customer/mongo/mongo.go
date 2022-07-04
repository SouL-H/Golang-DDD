package mongo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/percybolmer/ddd-go/aggregate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db       *mongo.Database
	customer *mongo.Collection
}

type mongoCustomer struct {
	ID   uuid.UUID `bson:"_id"`
	Name string    `bson:"name"`
}

func NewFromCustomer(c aggregate.Customer) mongoCustomer {
	return mongoCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}

func (m mongoCustomer) ToAggregate() aggregate.Customer {
	c := aggregate.Customer{}
	c.SetID(m.ID)
	c.SetName(m.Name)
	return c
}

func New(ctx context.Context, connection string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connection))
	if err != nil {
		return nil, err
	}
	db := client.Database("ddd")
	customers := db.Collection("customers")
	return &MongoRepository{
		db:       db,
		customer: customers,
	}, nil
}

func (mr *MongoRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := mr.customer.FindOne(ctx, bson.M{"_id": id})
	var c mongoCustomer
	err := result.Decode(&c)
	if err != nil {
		return aggregate.Customer{}, err
	}
	return c.ToAggregate(), nil
}

func (mr *MongoRepository) Add(c aggregate.Customer) error{
	ctx,cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	internal := NewFromCustomer(c)
	_, err := mr.customer.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}
func (mr *MongoRepository) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := mr.customer.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
func (mr *MongoRepository) Update(c aggregate.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	internal := NewFromCustomer(c)
	_, err := mr.customer.ReplaceOne(ctx, bson.M{"_id": c.GetID()}, internal)
	if err != nil {
		return err
	}
	return nil
}