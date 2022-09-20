package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "Black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Lee",
			"hair_color": "snow",
			"has_dog": false
		},
		{
			"first_name": "Adrew",
			"last_name": "Ng",
			"hair_color": "Black",
			"has_dog": true
		},
		{
			"first_name": "Charles",
			"last_name": "Wilson",
			"hair_color": "snow",
			"has_dog": false
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}
	log.Printf("unmashalled: %v", unmarshalled)

	// write a json from a struct
	var mySlice []Person
	var m1 Person
	m1.FirstName = "David"
	m1.LastName = "Beckham"
	m1.HairColor = "yellow"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Ana"
	m2.LastName = "Pytago"
	m2.HairColor = "yellow"
	m2.HasDog = true

	mySlice = append(mySlice, m2)

	var m3 Person
	m3.FirstName = "Taylor"
	m3.LastName = "Swift"
	m3.HairColor = "yellow"
	m3.HasDog = false

	mySlice = append(mySlice, m3)

	var m4 Person
	m4.FirstName = "King"
	m4.LastName = "Berbartop"
	m4.HairColor = "black"
	m4.HasDog = true

	mySlice = append(mySlice, m4)

	newJson, err := json.MarshalIndent(mySlice, "", "     ")
	if err != nil {
		log.Println("Error marshalling ", err)

	} else {
		log.Println("Json: ")
	}

	fmt.Println(string(newJson))

}
