package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter te rating  for our Pizza : ")

	// comma ok || err err

	// if there is an error the err will hold the error otherwise input will hold the variabl
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)

	fmt.Print("Please enter a number")
	number := bufio.NewReader(os.Stdin)
	number2, _ := number.ReadString('\n')
	fmt.Println(number2)
}
