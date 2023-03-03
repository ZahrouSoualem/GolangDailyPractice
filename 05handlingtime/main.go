package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello")
	start := time.Now()

	fmt.Println(start.Format("02-01-2006 Monday 15:4:5"))

	createdDate := time.Date(2020, time.August, 10, 10, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("02-01-2006 Monda 15:4:5"))

}
