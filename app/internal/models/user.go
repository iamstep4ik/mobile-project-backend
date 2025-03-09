package models

import "github.com/google/uuid"

type User struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Surname        string    `json:"surname"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
}

type UserUseCase interface {
	SignUp(user *User) (*User, error)
	Login(email, password string) (*User, error)
	Logout(id string) error
}

type UserRepository interface {
	CreateUser(signUp *User) (*User, error)
	CheckUser(login *User) (*User, error)
	GetByEmail(email string) (bool, error)
}
