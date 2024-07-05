package main

import (
	"fmt"
)

func main() {

	//1
	fmt.Println("//1")
	fmt.Println(upperCaseExcept([]string{"code", "java", "cool"}, "java"))
	fmt.Println("")
	fmt.Println("")
	
	//2
	fmt.Println("//2")
	fmt.Println(findMinMax([]int{2,3,4,5,6,7,8,9,20,10}))
	fmt.Println("")
	fmt.Println("")
	
	//3
	fmt.Println("//3")
	fmt.Println(findMinRange([]int{1, 22, 3, 4, 5, 10, 7, 8, 9, 49},0,7))
	fmt.Println("")
	fmt.Println("")

	fmt.Println("")
	fmt.Println(findMaxRange([]int{1, 22, 3, 4, 5, 10, 7, 8, 9, 49},2,7))
	fmt.Println("")
	fmt.Println("")

	//4 
	fmt.Println("//4")
	fmt.Println(evenOddOrder([]int{1,2,3,4,5,6,7,8,9,10}))
	fmt.Println("")
	fmt.Println("")
	
	//5
	fmt.Println("//5")
	rotateArray([]int{12,15,1,5,20},1)
	fmt.Println("")
	fmt.Println("")
	
	//6
	fmt.Println("//6")
	column1()
	fmt.Println("")
	fmt.Println("")
	
	//7
	fmt.Println("//7")
	column2()
	fmt.Println("")
	fmt.Println("")

	//8
	fmt.Println("//8")
	column3(7)
	fmt.Println("")
	fmt.Println("")
}
