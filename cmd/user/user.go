package user

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func MaybeAddUser(name string, userList *[]User) error {

	for _, user := range *userList {
		if user.Name == name {
			return errors.New("name already exists in session list")
		}
	}

	*userList = append(*userList, User{Name: name})
	return nil
}

func DeleteUser(userID uuid.UUID, userList *[]User) error {

	for i, user := range *userList {
		if user.ID == userID {
			uL := *userList
			uL[i] = uL[len(uL)-1]
			*userList = uL
			return nil
		}
	}

	return errors.New("user does not exist in session")
}
