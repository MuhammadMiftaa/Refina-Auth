package repository

import (
	"errors"

	"refina-auth/internal/types/model"

	"gorm.io/gorm"
)

type UsersRepository interface {
	GetAllUsers() ([]model.Users, error)
	GetUserByID(id string) (model.Users, error)
	GetUserByEmail(email string) (model.Users, error)
	CreateUser(user model.Users) (model.Users, error)
	UpdateUser(user model.Users) (model.Users, error)
	DeleteUser(user model.Users) (model.Users, error)
}

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &usersRepository{db}
}

func (user_repo *usersRepository) GetAllUsers() ([]model.Users, error) {
	var users []model.Users
	err := user_repo.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (user_repo *usersRepository) GetUserByID(id string) (model.Users, error) {
	var user model.Users
	err := user_repo.db.First(&user, "id = ?", id).Error
	if err != nil {
		return model.Users{}, errors.New("user not found")
	}

	return user, nil
}

func (user_repo *usersRepository) GetUserByEmail(email string) (model.Users, error) {
	var user model.Users
	err := user_repo.db.First(&user, "email = ?", email).Error
	if err != nil {
		return model.Users{}, errors.New("user not found")
	}

	return user, nil
}

func (user_repo *usersRepository) CreateUser(user model.Users) (model.Users, error) {
	err := user_repo.db.Create(&user).Error
	if err != nil {
		return model.Users{}, errors.New("failed to create user")
	}

	return user, nil
}

func (user_repo *usersRepository) UpdateUser(user model.Users) (model.Users, error) {
	err := user_repo.db.Save(&user).Error
	if err != nil {
		return model.Users{}, errors.New("failed to update user")
	}

	return user, nil
}

func (user_repo *usersRepository) DeleteUser(user model.Users) (model.Users, error) {
	err := user_repo.db.Delete(&user).Error
	if err != nil {
		return model.Users{}, errors.New("failed to delete user")
	}

	return user, nil
}
