package domain

type User struct{}

func (*User) HasAccess() bool {
	return true
}

func (*User) HomePath() string {
	return ""
}
