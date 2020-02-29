package entity

import "github.com/trewanek/go-echo-boiler/domain/value_object"

type User struct {
	userID   *value_object.UserID
	userName string
	age      int
	sex      string
}

func NewUser(userID *value_object.UserID, userName string, age int, sex string) *User {
	return &User{
		userID:   userID,
		userName: userName,
		age:      age,
		sex:      sex,
	}
}

func (u *User) UserID() *value_object.UserID {
	return u.userID
}

func (u *User) UserName() string {
	return u.userName
}

func (u *User) Age() int {
	return u.age
}

func (u *User) Sex() string {
	return u.sex
}

func (u *User) Equals(other *User) bool {
	if other == nil {
		return false
	}
	return u.userID.Equals(other.userID)
}
