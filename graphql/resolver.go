//go:generate go run github.com/99designs/gqlgen -v

package graphql

import (
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
