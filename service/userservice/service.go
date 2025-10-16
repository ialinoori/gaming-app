package userservice

import (
	"fmt"
	"gameapp/entity"
	"gameapp/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(phonenumber string) (bool, error)
	Reister(u entity.User) (entity.User, error)
}

type service struct {
	repo Repository
}

type RegisterRequest struct {
	Name        string
	PhoneNumber string
}

type RegisterResponse struct {
	User entity.User
}

func (s service) Register(req RegisterRequest) (RegisterResponse, error) {
	// validate phone number
	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number is not valid")
	}

	// check uniqueness of phone number
	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, err
		}
		return RegisterResponse{}, fmt.Errorf("phone number is not unique")
	}

	// validate name

	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf(("name length must be more than 3"))
	}

	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
	}

	// create new user in storage
	createdUser, err := s.repo.Reister(user)

	if err != nil {
		return RegisterResponse{}, fmt.Errorf(("failed to create new user"))
	}

	return RegisterResponse{User: createdUser}, nil
}
