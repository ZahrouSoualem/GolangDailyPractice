package main

import "fmt"

func main() {

	days := []string{"sunday", "Monday", "tuesday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Println(index, " and ", day)
	}

}
