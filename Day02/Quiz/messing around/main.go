package main

import "fmt"

func main() {
	a := 10

	fmt.Println("before pointer: ")
	fmt.Println(a)

	b := &a

	
	fmt.Println("printing b that hold adrress 'a' as value")
	fmt.Println(b)
	
	
	fmt.Println("printing b that refers to 'a' based on value address")
	fmt.Println(*b)
	
	*b = 2

	fmt.Println(a)
	fmt.Println(&a)

}