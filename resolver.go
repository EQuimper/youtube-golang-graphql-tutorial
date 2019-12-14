//go:generate go run github.com/99designs/gqlgen -v

package meetmeup

import (
	"context"
	"errors"

	"github.com/equimper/meetmeup/models"
	"github.com/equimper/meetmeup/postgres"
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

type Resolver struct {
	MeetupsRepo postgres.MeetupsRepo
	UsersRepo   postgres.UsersRepo
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
	return m.UsersRepo.GetUserByID(obj.UserID)
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
	if len(input.Name) < 3 {
		return nil, errors.New("name not long enough")
	}

	if len(input.Description) < 3 {
		return nil, errors.New("description not long enough")
	}

	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      "1",
	}

	return m.MeetupsRepo.CreateMeetup(meetup)
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return r.MeetupsRepo.GetMeetups()
}
