package graphql

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/equimper/meetmeup/middleware"
	"github.com/equimper/meetmeup/models"
)

var (
	ErrBadCredentials  = errors.New("email/password combination don't work")
	ErrUnauthenticated = errors.New("unauthenticated")
)

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (m *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	user, err := m.UsersRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, ErrBadCredentials
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return nil, ErrBadCredentials
	}

	token, err := user.GenToken()
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (m *mutationResolver) Register(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	_, err := m.UsersRepo.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("email already in used")
	}

	_, err = m.UsersRepo.GetUserByUsername(input.Username)
	if err == nil {
		return nil, errors.New("username already in used")
	}

	user := &models.User{
		Username:  input.Username,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err = user.HashPassword(input.Password)
	if err != nil {
		log.Printf("error while hashing password: %v", err)
		return nil, errors.New("something went wrong")
	}

	// TODO: create verification code

	tx, err := m.UsersRepo.DB.Begin()
	if err != nil {
		log.Printf("error creating a transaction: %v", err)
		return nil, errors.New("something went wrong")
	}
	defer tx.Rollback()

	if _, err := m.UsersRepo.CreateUser(tx, user); err != nil {
		log.Printf("error creating a user: %v", err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("error while commiting: %v", err)
		return nil, err
	}

	token, err := user.GenToken()
	if err != nil {
		log.Printf("error while generating the token: %v", err)
		return nil, errors.New("something went wrong")
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (m *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	meetup, err := m.MeetupsRepo.GetByID(id)
	if err != nil || meetup == nil {
		return false, errors.New("meetup not exist")
	}

	err = m.MeetupsRepo.Delete(meetup)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup: %v", err)
	}

	return true, nil
}

func (m *mutationResolver) UpdateMeetup(ctx context.Context, id string, input models.UpdateMeetup) (*models.Meetup, error) {
	meetup, err := m.MeetupsRepo.GetByID(id)
	if err != nil || meetup == nil {
		return nil, errors.New("meetup not exist")
	}

	didUpdate := false

	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("name is not long enough")
		}
		meetup.Name = *input.Name
		didUpdate = true
	}

	if input.Description != nil {
		if len(*input.Description) < 3 {
			return nil, errors.New("description is not long enough")
		}
		meetup.Description = *input.Description
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("no update done")
	}

	meetup, err = m.MeetupsRepo.Update(meetup)
	if err != nil {
		return nil, fmt.Errorf("error while updating meetup: %v", err)
	}

	return meetup, nil
}

func (m *mutationResolver) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}

	if len(input.Name) < 3 {
		return nil, errors.New("name not long enough")
	}

	if len(input.Description) < 3 {
		return nil, errors.New("description not long enough")
	}

	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      currentUser.ID,
	}

	return m.MeetupsRepo.CreateMeetup(meetup)
}
