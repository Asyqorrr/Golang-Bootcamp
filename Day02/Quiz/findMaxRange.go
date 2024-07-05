package main

import "sort"

func findMaxRange(arr []int, min int, max int) []int {
	result := []int{}
	valId := []int{}

	for i := min; i <= max; i++ {
		result = append(result, arr[i])
	}

	sort.Ints(result)

	small := result[len(result)-1]

	for i,v := range arr{
		if v == small {
			valId =append(valId, v,i)
			break
		}
	}
	return valId
}
