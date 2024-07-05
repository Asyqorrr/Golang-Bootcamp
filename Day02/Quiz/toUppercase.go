package main

import "strings"

func upperCaseExcept(arr []string, word string) []string {
	result := []string{}

	for _, v := range arr {
		if v == word {
			result = append(result, v)

		} else {
			result = append(result, strings.ToUpper(v))
		}

	}

	return result
}