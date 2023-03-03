package main

import "fmt"

func main() {

	//var one int = 2

	//var two = &one

	myNumber := 23

	var ptr = &myNumber

	fmt.Println(ptr)
	fmt.Println(*ptr)
}
