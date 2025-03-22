package procedures

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/profsmallpine/on-the-shelf/db/models"
	"github.com/profsmallpine/on-the-shelf/domain"
)

type User struct {
	allowedEmails []string
}

func (u *User) HandleCallback(data *domain.AuthData) (*domain.User, error) {
	if data.Email == "" || !u.isAllowedEmail(data.Email) {
		return nil, domain.ErrNotValid
	}

	user, err := queries.GetUserByEmail(context.Background(), data.Email)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}

		user, err = queries.CreateUser(context.Background(), models.CreateUserParams{
			Email:      data.Email,
			FirstName:  data.FirstName,
			LastName:   data.LastName,
			PictureUrl: data.PictureURL,
		})
		if err != nil {
			return nil, err
		}
	}

	return &domain.User{User: user}, nil
}

func (u *User) isAllowedEmail(email string) bool {
	for _, allowed := range u.allowedEmails {
		if email == allowed {
			return true
		}
	}
	return false
}
