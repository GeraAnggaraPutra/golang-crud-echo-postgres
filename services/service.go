package services

import (
	"crud-database-postgresql/models"
	"crud-database-postgresql/repository"
	"errors"
)

type Service interface {
	GetAll() ([]*models.User, error)
	FindById(id int) (*models.User, error)
	Create(user *models.User) error
	Update(id int, userUpdate *models.User) (*models.User, error)
	Delete(id int) error
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]*models.User, error) {
	users, err := s.repository.GetAll()
	return users, err
}

func (s *service) FindById(id int) (*models.User, error) {
	user, err := s.repository.FindById(id)
	return user, err
}

func (s *service) Create(user *models.User) error {
	err := s.repository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Update(id int, userUpdate *models.User) (*models.User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("User not found")
	}

	user.Name = userUpdate.Name
	user.Email = userUpdate.Email
	user.Age = userUpdate.Age
	user.Height = userUpdate.Height

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *service) Delete(id int) error {
	user, err := s.repository.FindById(id)
	if err != nil {
		return nil
	}

	if user == nil {
		return errors.New("user not found")
	}

	err = s.repository.Delete(user)
	if err != nil {
		return err
	}

	return nil
}