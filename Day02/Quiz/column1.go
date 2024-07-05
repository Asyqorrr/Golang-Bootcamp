package main

import "fmt"

func column1() {

	matrix := [5][5]int{}
	for r := 0; r <= 4; r++ {
		for c := 0; c <= 4; c++ {
			

			if r==c {
				matrix[r][c] = r+1
				fmt.Printf("%d ", matrix[r][c])

			} else if c > r {
				matrix[r][c] = 10
				fmt.Printf("%d ", matrix[r][c])
			
			} else {
				matrix[r][c] = 20
				fmt.Printf("%d ", matrix[r][c])
			}
		}

		fmt.Println("")
	}

}
