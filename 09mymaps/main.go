package main

import "fmt"

func main() {
	fmt.Println("Zahreddine Soualem")

	languages := make(map[string]string)

	languages["JS"] = "js"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages", languages)
	fmt.Println("List of all languages", languages["JS"])

	delete(languages, "RB")

	fmt.Println("List of all languages", languages)

	// iterate a language with a loop
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
	// a loop with comma ok
	for _, value := range languages {
		fmt.Printf(" value is %v\n", value)
	}

}
