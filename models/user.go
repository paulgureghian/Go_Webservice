package models

import (
	"errors"
	"fmt"
)

// User struct is exported
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers is exported
func GetUsers() []*User {

	return users
}

// AddUser is exported
func AddUser(u User) (User, error) {

	if u.ID != 0 {
		return User{}, errors.New("New user must not include id or it must be set to zero")
	}

	u.ID = nextID
	nextID++
	users = append(users, &u)

	return u, nil
}

// GetUserByID is exported
func GetUserByID(id int) (User, error) {

	for _, u := range users {

		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)

}

// UpdateUser is exported
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)

}

// RemoveUserByID is exported
func RemoveUserByID(id int) error {

	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '%v' not found \n", id)
}
