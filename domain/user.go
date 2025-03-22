package domain

import "github.com/profsmallpine/on-the-shelf/db/models"

type User struct {
	models.User
}

type UserProcedures interface {
	HandleCallback(data *AuthData) (*User, error)
}

func (*User) HasAccess() bool {
	return true
}

func (*User) HomePath() string {
	return "/books"
}
