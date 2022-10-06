package customer

import (
	"context"

	"developer.zopsmart.com/go/gofr/examples/data-layer-with-mongo/models"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type store struct{}

// New is factory function for store layer
//nolint:revive // customer should not be used without proper initilization with required dependency
func New() store {
	return store{}
}

// Get returns the list of models from mongodb based on the filter passed in the request
func (s store) Get(ctx *gofr.Context, name string) ([]models.Customer, error) {
	resp := make([]models.Customer, 0)

	// fetch the Mongo collection
	collection := ctx.MongoDB.Collection("customers")

	filter := bson.D{}

	if name != "" {
		nameFilter := primitive.E{
			Key:   "name",
			Value: name,
		}
		filter = append(filter, nameFilter)
	}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return resp, errors.DB{Err: err}
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var c models.Customer
		if err := cur.Decode(&c); err != nil {
			return resp, errors.DB{Err: err}
		}

		resp = append(resp, c)
	}

	return resp, nil
}

// Create extracts JSON content from request body and unmarshal it as Customer and then put it into db
func (s store) Create(ctx *gofr.Context, c models.Customer) error {
	// fetch the Mongo collection
	collection := ctx.MongoDB.Collection("customers")

	_, err := collection.InsertOne(context.TODO(), c)

	return err
}

// Delete deletes a record from MongoDB, returns delete count and the error if it fails to delete
func (s store) Delete(ctx *gofr.Context, name string) (int, error) {
	// fetch the Mongo collection
	collection := ctx.MongoDB.Collection("customers")
	filter := bson.D{}

	filter = append(filter, primitive.E{
		Key:   "name",
		Value: name,
	})

	deleted, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return 0, errors.DB{Err: err}
	}

	return int(deleted.DeletedCount), nil
}
