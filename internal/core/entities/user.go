package entities

import (
	"time"

	valueobjects "github.com/BrunoPolaski/auth-service/internal/core/value_objects"
)

type User struct {
	id                  int64
	email               valueobjects.Email
	password            valueobjects.Password
	needsPasswordChange bool
	createdAt           time.Time
	isActive            bool
}

func NewUser(
	email valueobjects.Email,
	password valueobjects.Password,
) User {
	return User{
		email:               email,
		password:            password,
		createdAt:           time.Now(),
		isActive:            true,
		needsPasswordChange: false,
	}
}

func NewUserWithFields(
	id int64,
	email valueobjects.Email,
	pwd valueobjects.Password,
	createdAt time.Time,
	isActive, needsPasswordChange bool,
) *User {
	return &User{
		id:                  id,
		email:               email,
		password:            pwd,
		needsPasswordChange: needsPasswordChange,
		createdAt:           createdAt,
		isActive:            isActive,
	}
}

func (u *User) Id() int64 {
	return u.id
}

func (u *User) Email() *valueobjects.Email {
	return &u.email
}

func (u *User) Password() *valueobjects.Password {
	return &u.password
}

func (u *User) NeedsPasswordChange() bool {
	return u.needsPasswordChange
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) IsActive() bool {
	return u.isActive
}

func (u *User) SetId(id int64) {
	u.id = id
}

func (u *User) SetEmail(email valueobjects.Email) {
	u.email = email
}

func (u *User) SetNeedsPasswordChange(needsPasswordChange bool) {
	u.needsPasswordChange = needsPasswordChange
}

func (u *User) SetCreatedAt(createdAt time.Time) {
	u.createdAt = createdAt
}

func (u *User) SetActive(isActive bool) {
	u.isActive = isActive
}
