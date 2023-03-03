package main

import "fmt"

func main() {

	loginCount := 23
	var result string
	if loginCount < 10 {
		result = " Regular user"
	} else if loginCount < 26 {
		result = " Regular user"
	} else {
		result = " Regular user"
	}

	fmt.Println(result)
}
