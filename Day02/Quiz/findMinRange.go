package main

import "sort"

func findMinRange(arr []int, min int, max int) []int {
	result := []int{}
	valId := []int{}

	for i := min; i <= max; i++ {
		result = append(result, arr[i])
	}

	sort.Ints(result)

	small := result[0]

	for i,v := range arr{
		if v == small {
			valId =append(valId, v,i)
			break
		}
	}
	return valId
}
