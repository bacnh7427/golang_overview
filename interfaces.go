package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shephered",
	}

	gorilla := Gorilla{
		Name:          "Gigdo",
		Color:         "Black Snow",
		NumberOfTeeth: 32,
	}

	PrintInfor(&dog)
	PrintInfor(&gorilla)
}

func PrintInfor(a Animal) {
	fmt.Println("This is animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "Aaaaaaaaaaaaaaaaaaaa"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}
