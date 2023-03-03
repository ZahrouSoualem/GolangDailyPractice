package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	OneAge int
}

func main() {
	fmt.Println("Welcome to struct")

	// no inheritance in golang: no super or parent
	hitesh := User{
		"zahrou",
		"zahrou001@gmail.Com",
		true,
		16,
		13,
	}

	fmt.Println(hitesh)
	fmt.Println(hitesh.Email)

	hitesh.GetStatus()

}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}
