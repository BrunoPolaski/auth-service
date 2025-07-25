package entities

import valueobjects "github.com/BrunoPolaski/auth-service/internal/core/value_objects"

type User struct {
	id                  int64
	email               string
	password            valueobjects.Password
	needsPasswordChange bool
	createdAt           string
	isActive            bool
}

func NewUser(email, password string) *User {
	pwd, err := valueobjects.NewPassword(password)

	return &User{
		email:    email,
		password: password,
	}
}

func (u *User) GetId() int64 {
	return u.id
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetNeedsPasswordChange() bool {
	return u.needsPasswordChange
}

func (u *User) GetCreatedAt() string {
	return u.createdAt
}

func (u *User) GetIsActive() bool {
	return u.isActive
}

func (u *User) SetId(id int64) {
	u.id = id
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) SetNeedsPasswordChange(needsPasswordChange bool) {
	u.needsPasswordChange = needsPasswordChange
}

func (u *User) SetCreatedAt(createdAt string) {
	u.createdAt = createdAt
}

func (u *User) SetActive(isActive bool) {
	u.isActive = isActive
}
