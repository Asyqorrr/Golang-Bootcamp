package main

import "fmt"

func column3(n int) {
	var matrix [][]int


	for r := 0; r < n; r++ {
		row := make([]int, 0, n)

		for c := 0; c < n; c++ {
			row = append(row, r)
		}
		fmt.Printf("%d ",row)
		fmt.Println("")

		matrix = append(matrix, row)
	}

}