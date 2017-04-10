package main

import "fmt"

const p string = "death and taxes"
const q = 42

func main() {
	const r = 45
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println("r - ", r)
}
