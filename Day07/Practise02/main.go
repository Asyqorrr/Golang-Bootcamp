package main

import (
	"day07/practise02/models"
	"fmt"
	"time"
)

func main() {
	
	emp1 := models.NewEmployee("Kang Dian", time.Now(), 10_000, models.Permanents)
	emp2 := models.NewEmployee("Yuli", time.Now(), 10_000, models.Permanents)
	emp3 := models.NewEmployee("Rini", time.Now(), 10_000, models.Permanents)

	//contract employee object
	contract1 := models.NewContract("Widi", time.Now(), 6_000, 150)

	//permanent employee object
	permanent1 := models.NewPermanent("Rizki", time.Now(), 10_000, 500)

	//slice of employee, context bisa diperlakukan tipe object yg berbeda
	employees := []models.Employee{*emp1, *emp2, *emp3, *&contract1.Employee, *&permanent1.Employee}

	for _, v := range employees {
		fmt.Println(v.ToString())
		fmt.Println("")
	}

	fmt.Println(contract1.Info())
}
