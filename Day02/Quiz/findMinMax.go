package main

import "sort"

func findMinMax(arr []int) []int {

	sort.Ints(arr)

	minMax := []int{
		arr[0],
		arr[len(arr)-1],
	}

	return minMax
}