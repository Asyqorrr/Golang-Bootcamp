package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02")
	fmt.Println("Current date:", currentDate)
}
