package userusecases

import (
	"context"
	domain "task_manager_api_clean_architecture/Domain"
	"task_manager_api_clean_architecture/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type userUsecaseSuite struct {
	suite.Suite
	repository *mocks.UserRepository

	usecase domain.UserUseCase
}


func (suite *userUsecaseSuite) SetupTest() {
	repository := new(mocks.UserRepository)
	usecase := NewUserUseCase(repository, 10 * time.Second)

	suite.repository = repository
	suite.usecase = usecase
}

func (suite *userUsecaseSuite) TestRegister() {
	user := &domain.User{
		Email: "test@example.com",
		Password: "password123",
	}

	suite.repository.On("Create", mock.Anything, user).Return(nil)

	err := suite.repository.Create(context.Background(), user)

	suite.NoError(err)
	suite.repository.AssertExpectations(suite.T())
}

func (suite *userUsecaseSuite) TestUserLogin_Positive() {
	user := domain.User{
		Email:    "test@example.com",
		Password: "password123",
	}

	expectedToken := "mocked-jwt-token"

	suite.repository.On("Login", mock.Anything, &user).Return(expectedToken, nil)

	token, err := suite.usecase.Login(context.Background(), &user)

	// Assertions
	suite.NoError(err)
	suite.Equal(expectedToken, token)
	suite.repository.AssertExpectations(suite.T())
}

func (suite *userUsecaseSuite) TestPromote() {
	userId := "some-user-id"

	suite.repository.On("Promote", mock.Anything, userId).Return(nil)

	// Call the method being tested
	err := suite.usecase.Promote(context.Background(), userId)

	// Assertions
	suite.NoError(err)
	suite.repository.AssertExpectations(suite.T())
}

func TestUserUsecase(t *testing.T) {
	suite.Run(t, new(userUsecaseSuite))
}



