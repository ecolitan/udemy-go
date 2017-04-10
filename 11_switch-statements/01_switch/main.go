package main

import "fmt"

func main() {
	i := "a"
	switch i {
	case "Daniel", "a":
		fmt.Println("Hi Daniel")
	case "Medhi":
		fmt.Println("Hi Medhi")
	case "Jenny":
		fmt.Println("Hi Jenny")
	default:
		fmt.Println("Have you no friends?")
	}
}
