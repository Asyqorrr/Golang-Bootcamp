package main

import (
	"fmt"
	s "strings"
)

func main() {

	// 1. findDivisor
	findDivisor(7)

	// 2. reverseInt
	reverseInt(5432)

	// 3. Triangle 1
	Triangle1(4)

	// 4. Triangle2
	Triangle2(4)

	// 5. Deret
	Deret(5)

	// 6. Deret Selang-Seling
	SelangSeling(5)

	//8 Palindrome
	fmt.Println(isPalindrome("aku Usa"))
	
	//9 Reverse
	fmt.Println(reverse("XYnb"))
	fmt.Println("-----------")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	//10 CheckBraces
	fmt.Println(checkBraces("()))((()"))
	fmt.Println("-----------")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	//11 NumPalindrome
	fmt.Print(isNumberPalindrome(12221))
}

// 1.
func findDivisor(n int) {
	for i := 1; i < n; i++ {

		if n%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
}

// 2.
func reverseInt(n int){
	sisa := n

	for{
		digit := sisa % 10
		fmt.Printf("%d ",digit)

		sisa = sisa/10

		if sisa == 0{
			break
		}
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
}

// 3.
func Triangle1 (n int){
	for row := 0; row < n ; row++{
		for col := 0; col <= n; col++{
			
			if row >= col {
				fmt.Print("   ")
			
			}else{
				fmt.Print(" * ")
			}
			
		}

		fmt.Println("")
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
}

// 4.
func Triangle2 (n int){
	for row := 0; row < n ; row++{
		for col := 0; col <= n; col++{
			
			if row + col >= 4{
				fmt.Print(" * ")
			}else{
				fmt.Print("   ")
			}
		}

		fmt.Println("")
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
}

// 5.
func Deret (n int){
	
	firstN:= 1
	secondN:= n

	increaser := 0
	// decreaser := 0

	for r:= 1; r <=n; r++{
		for c:= 1; c <= n/2; c++{
			fmt.Printf("%d ",firstN+increaser)
			fmt.Printf("%d ",secondN-increaser)
		}
		increaser++
		fmt.Println("")
	}
	fmt.Println("----------------")
	fmt.Println("")
	fmt.Println("")
}

// 6.
func SelangSeling (n int){
	for row := 1; row <= n; row++{
		for col := 1; col <= n; col++{
			
			if (row+col) % 2 == 0{
				fmt.Print("-")
			}else{
				fmt.Printf("%d",col)
			}
		}
		fmt.Println("")
	} 
	// fmt.Println("-------------------")
	// fmt.Println("")
	// fmt.Println("")
}

// 8.
func isPalindrome(word string)bool{
	word = s.ToLower(word)
	
	reversed := ""

	for i:= len(word)-1; i >= 0; i--{
		reversed += string(rune(word[i])) 

	}
	fmt.Println("-----------")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	return word == reversed
}

// 9.
func reverse (word string)string{
	
	reversed := ""

	for i:= len(word)-1; i >= 0; i--{
		reversed += string(rune(word[i])) 

	}
	fmt.Println("-----------")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	return reversed
}

func checkBraces(s string)bool{
	count:= 0

	for _, char := range s{
		if char == '('{
			count++
		
		}else if char == ')'{
			if count == 0{
				return false
			}
			count--
		}
	}

	return count == 0
}

func isNumberPalindrome(n int)bool{
	reversed := 0
	check := n
	sisa := n 
	for {
		digit := sisa % 10
		reversed = reversed * 10 + digit
		sisa = sisa / 10

		if sisa == 0{
			
			if reversed == check{
				return true

			}else {
				return false
			}
		}
	}
}