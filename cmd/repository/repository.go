package repository

import (
	"context"
	"errors"

	"github.com/PorcoGalliard/rumahweb-interview/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	Database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Database: db,
	}
}


func (r *UserRepository) CreateNewUser(ctx context.Context, user *models.User) (int64, error) {
	err := r.Database.WithContext(ctx).Create(user).Error
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.Database.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByUserID(ctx context.Context, userID int64) (*models.User, error) {
	var user models.User
	err := r.Database.WithContext(ctx).Where("id = ?", userID).Last(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := r.Database.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	err := r.Database.WithContext(ctx).Table("users").Save(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
} 

func (r *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	err := r.Database.WithContext(ctx).Table("users").Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}