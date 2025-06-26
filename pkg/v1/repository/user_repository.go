// Package repository provides the implementation of the repository interfaces.
package repository

import (
	"context"

	"github.com/dawit_hopes/grpc_micro_service/internal/domain/models"
	interfaces "github.com/dawit_hopes/grpc_micro_service/pkg/v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db  *mongo.Collection
	ctx context.Context
}

func NewUserRepository(db *mongo.Collection) interfaces.RepoInterface {
	return &userRepository{db: db, ctx: context.Background()}
}

func (r *userRepository) Create(user models.User) error {
	_, err := r.db.InsertOne(r.ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Get(id string) (models.User, error) {
	var user models.User
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	err = r.db.FindOne(r.ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Update(user models.User) error {
	objID, err := primitive.ObjectIDFromHex(user.ID.Hex())
	if err != nil {
		return err
	}
	_, err = r.db.UpdateOne(r.ctx, bson.M{"_id": objID}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.db.DeleteOne(r.ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}
