package main

import "fmt"

func main() {
	fmt.Println("Welcome to function in golang")
	result := adder(1, 3)
	fmt.Println(result)
	greeter()
	result2, result3 := proAdd(1, 2, 3, 4)
	fmt.Println(result2, "-", result3)
}

func greeter() {
	fmt.Println("Nmastey from golang")
}

func adder(a int, b int) int {
	return a + b
}

func proAdd(values ...int) (int, int) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, 1
}
