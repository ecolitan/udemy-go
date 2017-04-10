package main

import "fmt"

func main() {
	var m int
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	fmt.Print("Enter a larger number: ")
	fmt.Scan(&m)
	fmt.Println(m % n)
}
