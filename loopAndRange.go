package main

import "log"

func main() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
		Mobile    string
	}

	var users []User
	users = append(users, User{"John", "Smith", "johnsmith@gmail.com", 39, "123121212"})
	users = append(users, User{"Harry", "Potter", "Potter@gmail.com", 35, "123121234"})
	users = append(users, User{"King", "David", "King@gmail.com", 45, "1231212345"})
	users = append(users, User{"Chales", "Smith", "Chales@gmail.com", 30, "1231212346"})
	users = append(users, User{"Eliz", "Smith", "Eliz@gmail.com", 40, "1231212127"})
	users = append(users, User{"John", "Biden", "John@gmail.com", 50, "1231212128"})
	users = append(users, User{"Donald", "Trump", "Donald@gmail.com", 20, "1231212912"})
	users = append(users, User{"Brack", "Obama", "Brack@gmail.com", 50, "12318212182"})
	users = append(users, User{"Pens", "Puck", "Pens@gmail.com", 90, "12312121892"})
	users = append(users, User{"Alaa", "Toleiv", "Toleiv@gmail.com", 80, "12312891212"})
	users = append(users, User{"Beckham", "David", "Beckham@gmail.com", 70, "12389121212"})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age, l.Mobile)
	}
}
