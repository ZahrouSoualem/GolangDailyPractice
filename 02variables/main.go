package main

import "fmt"

const LongToken string = "ghbbhhjd"

func main() {
	fmt.Println("Variables")

	var username string = "hitech"
	fmt.Println(username)
	fmt.Printf("variable if type of : %T \n ", username)

	var isLogged bool = true
	fmt.Println(isLogged)
	fmt.Printf("variable if type of : %T \n ", isLogged)

	var smallFloat float64 = 123.3447865875
	fmt.Println(smallFloat)
	fmt.Printf("variable if type of : %09.2f \n ", smallFloat)

	var smallvariabble int
	fmt.Println(smallvariabble)
	fmt.Printf("variable if type of : %T \n ", smallvariabble)

	var smallvariabble2 = "How are you"
	fmt.Println(smallvariabble2)

	// no var style

	variable := 300000
	fmt.Println(variable)

	fmt.Println(LongToken)
	fmt.Printf("variable if type of : %T \n ", LongToken)
}
