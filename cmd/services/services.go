package services

import (
	"context"
	"errors"
	"time"

	"github.com/PorcoGalliard/rumahweb-interview/cmd/repository"
	"github.com/PorcoGalliard/rumahweb-interview/models"
	"github.com/PorcoGalliard/rumahweb-interview/utils"
	"github.com/golang-jwt/jwt/v5"
)

type UserServices struct {
	UserRepo repository.UserRepository
	JWTSecret string
}

func NewUserServices(userRepo *repository.UserRepository, JWTSecret string) *UserServices {
	return &UserServices{
		UserRepo: *userRepo,
		JWTSecret: JWTSecret,
	}
}

func (s *UserServices) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.UserRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServices) GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	user, err := s.UserRepo.FindByUserID(ctx, id)
	if err != nil {
		return nil,err
	}

	return user, nil
}

func (s *UserServices) GetAllUser(ctx context.Context) (*[]models.User, error) {
	users, err := s.UserRepo.FindAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (s *UserServices) CreateNewUser(ctx context.Context, user *models.User) (int64, error) {
	userID, err := s.UserRepo.CreateNewUser(ctx, user)
	if err != nil {
		return 0, nil
	}

	return userID, nil
}

func (s *UserServices) RegisterUser (ctx context.Context, user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	_, err = s.CreateNewUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserServices) LoginUser (ctx context.Context, params *models.LoginParameter) (string, error) {
	user, err := s.GetUserByEmail(ctx, params.Email)
	if err != nil {
		return "", err
	}

	isMatch, err := utils.CheckPasswordHash(user.Password, params.Password)
	if err != nil {
		return "", err
	}

	if !isMatch {
		return "", errors.New("Invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *UserServices) DeleteUser (ctx context.Context, id int64) error {
	if err := s.DeleteUser(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *UserServices) UpdateUser (ctx context.Context, user *models.User) (*models.User, error) {
	updatedUser, err := s.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}