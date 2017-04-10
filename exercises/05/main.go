package main

import "fmt"

func main() {
	var i int
	for i = 0; i < 101; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
