package auth

import (
	"context"
	"fmt"
	"time"

	"go.boilerplate/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	dbUser *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{dbUser: db.Collection("users")}
}

func (r *Repository) SignUpRepository(ctx context.Context, name string, email string, password string) (*model.User, error) {

	user := model.User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	resp, err := r.dbUser.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = resp.InsertedID.(primitive.ObjectID)

	return &user, nil
}

func (r *Repository) FindUserRepository(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	filter := bson.M{"email": email}
	err := r.dbUser.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, err
	}

	return &user, nil

}

func (r *Repository) ResetPasswordRepository(ctx context.Context, id string, password string) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, fmt.Errorf("invalid ID format: %v", err)
	}

	filter := bson.M{"_id": objectID}

	update := bson.M{
		"$set": bson.M{
			"password":   password,
			"updated_at": time.Now(),
		},
	}

	result, err := r.dbUser.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, fmt.Errorf("failed to update password: %v", err)
	}

	if result.MatchedCount == 0 {
		return false, fmt.Errorf("no user found with ID %s", id)
	}

	return true, nil
}
