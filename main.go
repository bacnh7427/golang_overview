package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, see you again"
	i = 10

	fmt.Println("Hi. " + whatToSay)
	fmt.Println("Hi. ", i)

	whatWasSay, els := saySomething()
	fmt.Println(whatWasSay, els)
}

func saySomething() (string, string) {
	return "Something", "Else"
}
