package userrepository

import (
	"context"
	"errors"
	"fmt"
	"os"
	domain "task_manager_api_clean_architecture/Domain"
	infrastructure "task_manager_api_clean_architecture/Infrastructure"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	database *mongo.Database
	collection string
}

func NewUserRepository(db *mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database: db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	hashedPassword, err := infrastructure.GeneratePasswordHash(user.Password)
	if err != nil {
		return errors.New("internal server error: failed to hash password")
	}

	collection := ur.database.Collection(ur.collection)
	count, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return errors.New("internal server error: failed to count documents")
	}

	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	filter := bson.M{"email": user.Email}
	err = collection.FindOne(context.Background(), filter).Err()

	if err == mongo.ErrNoDocuments {
		user.Password = hashedPassword
		_, insertErr := collection.InsertOne(context.Background(), user)
		if insertErr != nil {
			return errors.New("internal server error: failed to insert user")
		}
		return nil
	}

	if err == nil {
		return fmt.Errorf("user already exists")
	}

	return errors.New("internal server error: failed to check if user exists")
}

func (ur *userRepository) Login(c context.Context, user *domain.User) (string, error) {
	filter := bson.M{"email": user.Email}
	var existingUser domain.User

	collection := ur.database.Collection(ur.collection)
	err := collection.FindOne(context.Background(), filter).Decode(&existingUser)
	if err != nil {
		return "", err
	}

	var jwtSecret = []byte(getJwtSecret("JWT_SECRET"))

	isSame := infrastructure.ComparePasswordHash(user.Password, existingUser.Password)

	if !isSame {
		return "", errors.New("invalid email or password")
	}

	jwtService := infrastructure.NewJWTService(jwtSecret)
	jwtToken, err := jwtService.GenerateToken(existingUser.ID.String(), existingUser.Email, existingUser.Role, 24 * 60 * 60 * 30)

	if err != nil {
		return "", errors.New("inernal serever error")
	}

	return jwtToken, nil
}

func (ur *userRepository) Promote(c context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}

	collection := ur.database.Collection(ur.collection)
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"role": "admin"}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}

func getJwtSecret(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		return err.Error()
	}

	return os.Getenv(key)
}