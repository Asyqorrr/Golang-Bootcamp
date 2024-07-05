package main

import "fmt"

func rotateArray(arr []int, n int) {
	
	// first := arr[0]

	// arr[0] = arr[1]
	// arr[1] = arr[2]
	// arr[2] = arr[3]
	// arr[3] = arr[4]
	// arr[4] = first


	for i:=1; i <= n; i++ {
		first := arr[0]

		for i := range arr{
			if i == len(arr)-1{
				arr[i] = first

			}else{
				arr[i] = arr[i+1]
			}
		}
	}

	fmt.Println(arr)
}
