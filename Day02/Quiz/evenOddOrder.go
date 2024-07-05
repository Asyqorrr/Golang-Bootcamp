package main

func evenOddOrder(arr []int) []int {

	odd := []int{}
	even := []int{}

	for _, v := range arr {
		if v%2 == 0 {
			even = append(even, v)

		} else {
			odd = append(odd, v)
		}
	}

	even = append(even, odd...)

	return even
}