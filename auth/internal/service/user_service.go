package service

import (
	"log"

	"github.com/Solutions-Corp/chetapp-backend/auth/internal/config"
	"github.com/Solutions-Corp/chetapp-backend/auth/internal/model"
	"github.com/Solutions-Corp/chetapp-backend/auth/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateDefaultUser() error
	GetUserByEmail(email string) (*model.User, error)
	GetUserByID(id uuid.UUID) (*model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) CreateDefaultUser() error {
	config, err := config.LoadConfig()
	if err != nil {
		return err
	}

	existingUser, err := s.userRepository.GetUserByEmail(config.DefaultUserEmail)
	if err != nil {
		return err
	}

	if existingUser != nil {
		log.Println("Default user already exists")
		return nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(config.DefaultUserPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	defaultUser := &model.User{
		Email:        config.DefaultUserEmail,
		PasswordHash: string(hashedPassword),
		FirstName:    "CHETAPP",
		LastName:     "ADMINISTRATOR",
	}

	log.Println("Creating default user:", config.DefaultUserEmail)
	return s.userRepository.CreateUser(defaultUser)
}

func (s *userService) GetUserByEmail(email string) (*model.User, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByID(id uuid.UUID) (*model.User, error) {
	return s.userRepository.GetUserByID(id)
}
