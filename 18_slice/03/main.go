package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5)

	fmt.Println("-------------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("-------------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value:", mySlice[i])
	}

	testSlice := make([]interface{}, 0, 5)
	testSlice = append(testSlice, 1)
	testSlice = append(testSlice, "Hello, World!")
	fmt.Printf("%T\n", testSlice[0])
	fmt.Printf("%T\n", testSlice[1])

}
