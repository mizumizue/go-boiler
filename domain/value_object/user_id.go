package value_object

import "github.com/google/uuid"

type UserID struct {
	value string
}

func NewUserID() *UserID {
	return &UserID{value: generateID()}
}

func generateID() string {
	return uuid.New().String()
}

func (u *UserID) Value() string {
	return u.value
}

func (u *UserID) Equals(other *UserID) bool {
	if other == nil {
		return false
	}
	return u.value == other.value
}
