package userrepository

import (
	"context"
	"fmt"
	config "task_manager_api_clean_architecture/Config"
	domain "task_manager_api_clean_architecture/Domain"
	"task_manager_api_clean_architecture/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositorySuite struct {
	suite.Suite
	repository domain.UserRepository
	db         *mongo.Database
	jwtService *mocks.JWTService
}

func (suite *userRepositorySuite) SetupSuite() {
	configs, err := config.LoadConfig("../../")
	fmt.Println("view confg", configs)
	suite.NoError(err)
	suite.db = config.GetDB(&configs)
	suite.jwtService = new(mocks.JWTService)
	repository := NewUserRepository(suite.db, "users", suite.jwtService)

	suite.repository = repository
}

func (suite *userRepositorySuite) TearDownSuite() {
	err := suite.db.Drop(context.Background())
	suite.NoError(err)
}

func (suite *userRepositorySuite) TearDownTest() {
	err := suite.db.Collection("users").Drop(context.Background())
	suite.NoError(err)

}

func (suite *userRepositorySuite) TestCreateUser_Positive() {
	user := &domain.User{
		Email: "test@example.com",
		Password: "password123",
	}

	err := suite.repository.Create(context.Background(), user)
	suite.NoError(err)
}

func (suite *userRepositorySuite) TestCreateUser_Duplicate() {
	user := &domain.User{
		Email: "test@example.com",
		Password: "password123",
	}

	err := suite.repository.Create(context.Background(), user)
	suite.NoError(err)

	err = suite.repository.Create(context.Background(), user)
	suite.Error(err)
}

func (suite *userRepositorySuite) TestLogin_Positive() {
	user := &domain.User{
		Email:    "test@example.com",
		Password: "password123",
	}

	err := suite.repository.Create(context.Background(), user)
	suite.NoError(err)
	user.Password = "password123"

	suite.jwtService.On("GenerateToken", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("mocked_token", nil)

	jwtToken, err := suite.repository.Login(context.Background(), user)
	suite.NoError(err)
	suite.NotEmpty(jwtToken)
}

func (suite *userRepositorySuite) TestLogin_Invalid() {
	user := &domain.User{
		Email:    "test@example.com",
		Password: "wrongpassword",
	}


	_ = suite.repository.Create(context.Background(), &domain.User{
		Email:    "test@example.com",
		Password: "password123",
	})

	jwtToken, err := suite.repository.Login(context.Background(), user)
	suite.Error(err)
	suite.Empty(jwtToken)
}

func (suite *userRepositorySuite) TestPromoteUser_Positive() {
	user1 := &domain.User{
		Email:    "user1@example.com",
		Password: "password123",
	}

	user2 := &domain.User{
		Email:    "user2@example.com",
		Password: "password123",
	}

	// Create the user
	err := suite.repository.Create(context.Background(), user1)
	suite.NoError(err)

	err = suite.repository.Create(context.Background(), user2)
	suite.NoError(err)

	// Promote the user
	err = suite.repository.Promote(context.Background(), user2.ID.Hex())
	suite.NoError(err)
}

func (suite *userRepositorySuite) TestPromoteUser_InvalidID() {
	err := suite.repository.Promote(context.Background(), "invalidID")
	suite.Error(err)
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(userRepositorySuite))
}