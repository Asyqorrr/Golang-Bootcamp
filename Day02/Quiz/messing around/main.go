package main

import "fmt"

func main() {

	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Eve":   28,
	}

	ages["Alice"]++

	fmt.Println(ages)
}