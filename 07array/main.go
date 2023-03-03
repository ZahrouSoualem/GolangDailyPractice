package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[1] = "A"
	fruitList[2] = "C"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList = [3]string{"A", "B", "C"}
	fmt.Println(vegList)

}
