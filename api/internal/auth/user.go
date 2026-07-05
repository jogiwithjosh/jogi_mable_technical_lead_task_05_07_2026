package auth

import "time"

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"createdAt"`
}

func (u User) Public() PublicUser {
	return PublicUser{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

type PublicUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
