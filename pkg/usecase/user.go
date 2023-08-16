package usecase

import (
	"sha/pkg/domain"
	"sha/pkg/repository"

)

type UserUseCase struct {
	UserRepo repository.UserRepository
}

func (u *UserUseCase) GetAllUsers() ([]domain.User, error) {
	return u.UserRepo.GetAllUsers()
}
