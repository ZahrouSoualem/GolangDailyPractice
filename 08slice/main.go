package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Println(fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highStores := make([]int, 4)

	highStores[0] = 234
	highStores[1] = 345
	highStores[2] = 234
	highStores[3] = 345

	highStores = append(highStores, 555, 666, 777)
	fmt.Println(highStores)
	sort.Ints(highStores)
	fmt.Println(highStores)

	courses := []string{"reactjs", "javascript", "swift", "python", "ruby"}
	inedx := 2
	courses = append(courses[:inedx], courses[inedx+1:]...)
	fmt.Println(courses)
}
