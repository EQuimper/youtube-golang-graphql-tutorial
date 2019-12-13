//go:generate go run github.com/99designs/gqlgen -v

package meetmeup

import (
	"context"
	"errors"

	"github.com/equimper/meetmeup/models"
)

var meetups = []*models.Meetup{
	{
		ID:          "1",
		Name:        "A meetup",
		Description: "A description",
		UserID:      "1",
	},
	{
		ID:          "2",
		Name:        "A second meetup",
		Description: "A description",
		UserID:      "2",
	},
}

var users = []*models.User{
	{
		ID:       "1",
		Username: "Bob",
		Email:    "bob@gmail.com",
	},
	{
		ID:       "2",
		Username: "Jon",
		Email:    "jon@gmail.com",
	},
}

type Resolver struct {
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

type meetupResolver struct{ *Resolver }

func (m *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	user := new(models.User)

	for _, u := range users {
		if u.ID == obj.UserID {
			user = u
			break
		}
	}

	if user == nil {
		return nil, errors.New("user with id not exist")
	}

	return user, nil
}

func (r *Resolver) Meetup() MeetupResolver {
	return &meetupResolver{r}
}

type userResolver struct{ *Resolver }

func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

func (u *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	var m []*models.Meetup

	for _, meetup := range meetups {
		if meetup.UserID == obj.ID {
			m = append(m, meetup)
		}
	}

	return m, nil
}

func (m *mutationResolver) CreateMeetup(ctx context.Context, input NewMeetup) (*models.Meetup, error) {
	panic("implement me")
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return meetups, nil
}
