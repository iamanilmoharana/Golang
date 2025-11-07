package main

import "fmt"

const age = 30

func main() {
	const name = "golang"

	// name = "java" --cannot assign to name

	// fmt.Println(age)

	//Grouping
	const (
		port = 500
		host = "localhost"
	)

	fmt.Println(port, host)
}