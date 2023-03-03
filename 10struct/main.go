package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Welcome to struct")

	// no inheritance in golang: no super or parent
	hitesh := User{
		"zahrou",
		"zahrou001@gmail.Com",
		true,
		16,
	}

	fmt.Println(hitesh)
	fmt.Println(hitesh.Email)
}
