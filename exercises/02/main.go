package main

import "fmt"

func main() {
	var v string
	fmt.Printf("Enter your name: ")
	fmt.Scan(&v)
	fmt.Println("Hello,", v)
}
