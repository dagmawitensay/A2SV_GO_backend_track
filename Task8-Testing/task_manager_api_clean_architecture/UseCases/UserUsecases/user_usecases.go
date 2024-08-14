package userusecases

import (
	"context"
	domain "task_manager_api_clean_architecture/Domain"
	"time"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUseCase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUseCase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (uu *userUsecase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.userRepository.Create(ctx, user)
}

func (uu *userUsecase) Login(c context.Context, user *domain.User)(string, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.userRepository.Login(ctx, user)
}

func (uu *userUsecase) Promote(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.userRepository.Promote(ctx, id)
}