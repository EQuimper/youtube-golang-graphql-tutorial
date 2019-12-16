//go:generate go run github.com/99designs/gqlgen -v

package graphql

import (
	"github.com/equimper/meetmeup/postgres"
)

type Resolver struct {
	MeetupsRepo postgres.MeetupsRepo
	UsersRepo   postgres.UsersRepo
}
