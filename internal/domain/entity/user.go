package entity

type User struct {
	Id                  int64  `json:"id,omitempty"`
	Email               string `json:"email"`
	Password            string `json:"password"`
	NeedsPasswordChange bool   `json:"needs_password_change,omitempty"`
	CreatedAt           string `json:"created_at"`
	IsActive            bool   `json:"is_active"`
}

func NewUser(email, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

func (u *User) GetId() int64 {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetNeedsPasswordChange() bool {
	return u.NeedsPasswordChange
}

func (u *User) GetCreatedAt() string {
	return u.CreatedAt
}

func (u *User) GetIsActive() bool {
	return u.IsActive
}

func (u *User) SetId(id int64) {
	u.Id = id
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) SetNeedsPasswordChange(needsPasswordChange bool) {
	u.NeedsPasswordChange = needsPasswordChange
}

func (u *User) SetCreatedAt(createdAt string) {
	u.CreatedAt = createdAt
}

func (u *User) SetActive(isActive bool) {
	u.IsActive = isActive
}
