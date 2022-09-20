package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	Birthday    time.Time
}

func main() {
	user := User{
		FirstName:   "Bac",
		LastName:    "Nguyen",
		PhoneNumber: "123-456-789",
	}

	log.Println(user.FirstName, user.LastName, user.PhoneNumber, "; Birthday: ", user.Birthday)
}
