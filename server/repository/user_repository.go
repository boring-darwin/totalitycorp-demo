package repository

import (
	"errors"
	"log"
)

type UserItem struct {
	Id      int32
	Fname   string
	City    string
	Phone   int32
	Height  float32
	Married bool
}

var userList []*UserItem

func InitMockDB() {

	u1 := &UserItem{
		Id:      1,
		Fname:   "Bruce Wayne",
		City:    "Gautham",
		Phone:   996574634,
		Height:  6.2,
		Married: false,
	}
	userList = append(userList, u1)

	u2 := &UserItem{
		Id:      2,
		Fname:   "Peter Parker",
		City:    "Gautham",
		Phone:   996574634,
		Height:  5.7,
		Married: false,
	}
	userList = append(userList, u2)

	u3 := &UserItem{
		Id:      3,
		Fname:   "Matt Murdock",
		City:    "New york",
		Phone:   996574634,
		Height:  6.1,
		Married: false,
	}
	userList = append(userList, u3)

	log.Println("Initialized MockDB")

}

func (user UserItem) insertUser() {

}

func FindUser(id int32) (UserItem, error) {

	for _, i := range userList {

		if i.Id == id {
			return *i, nil
		}
	}

	return UserItem{}, errors.New("not found")
}

func (user UserItem) removeUser() {

}
