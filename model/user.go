package model

import (
	"errors"
	"fmt"
)

// User example
type User struct {
	ID   string       `json:"id" example:"1"`
	Name string    `json:"name" example:"user name"`
}

// AddUser example
type AddUser struct {
	ID   string       `json:"id" example:"1"`
	Name string `json:"name" example:"user name"`
}

var (
	ErrNameInvalid = errors.New("name is empty")
	ErrNoRow = errors.New("no row")
)

// Validation example
func (a AddUser) Validation() error {
	switch {
	case len(a.Name) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

// UpdateUser example
type UpdateUser struct {
	Name string `json:"name" example:"user name"`
}

// Validation example
func (a UpdateUser) Validation() error {
	switch {
	case len(a.Name) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

// UsersAll example
func UsersAll(q string) ([]User, error) {
	if q == "" {
		return users, nil
	}
	as := []User{}
	for k, v := range users {
		if q == v.Name {
			as = append(as, users[k])
		}
	}
	return as, nil
}

// UserOne example
func UserOne(id string) (User, error) {
	for _, v := range users {
		if id == v.ID {
			return v, nil
		}
	}
	return User{}, ErrNoRow
}

// Insert example
func (a User) Insert() error {
	users = append(users, a)
	return nil
}

// DeleteUser example
func DeleteUser(id string) error {
	for k, v := range users {
		if id == v.ID {
			users = append(users[:k], users[k+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user id=%q is not found", id)
}

// Update example
func (a User) Update() error {
	for k, v := range users {
		if a.ID == v.ID {
			users[k].Name = a.Name
			return nil
		}
	}
	return fmt.Errorf("user id=%q is not found", a.ID)
}

var users = []User{
	{ID: "1", Name: "user_1"},
	{ID: "2", Name: "user_2"},
	{ID: "3", Name: "user_3"},
}
