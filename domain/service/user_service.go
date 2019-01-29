package service

import (
	"fmt"

	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/domain/repository"
)

type UserService interface {
	RegisterUser(email, password string) (*model.User, error)
	ComparePasswords(plain, hash string) error
}

type userService struct {
	repository  repository.UserRepository
	hashService HashService
}

func NewUserService(repo repository.UserRepository, hashService HashService) UserService {
	return &userService{
		repository:  repo,
		hashService: hashService,
	}
}

func (s userService) RegisterUser(email, password string) (*model.User, error) {
	if err := s.checkIfEmailIsUnique(email); err != nil {
		return nil, err
	}

	hash, err := s.hashService.GenerateHashFromPassword(password)
	if err != nil {
		return nil, err
	}

	return model.NewUser(email, hash), nil
}

func (s userService) checkIfEmailIsUnique(email string) error {
	user, _ := s.repository.FindByEmail(email)
	if user != nil {
		return fmt.Errorf("email is not unique: %s", email)
	}

	return nil
}

func (s userService) ComparePasswords(plain, hash string) error {
	return s.hashService.ComparePasswords(plain, hash)
}
