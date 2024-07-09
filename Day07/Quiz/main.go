package main

import (
	models "day07/practise02/quiz/models"
	"fmt"
	"time"
)

func main() {
	currentTime:= time.Now()
	currentDate := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	
	
	// emp1 := models.NewEmployee("Asyqor", "Thoriq", currentDate,  models.Programmer, 6_000_000)
	// emp2 := models.NewEmployee("Riki", "Raka", currentDate,  models.Sales, 3_000_000)
	// emp3 := models.NewEmployee("Lord", "Stark", currentDate,  models.Sales, 3_000_000)
	// emp4 := models.NewEmployee("Mike", "Radahn", currentDate,  models.QA, 4_500_000)
	// emp5 := models.NewEmployee("Miquella", "Malika", currentDate,  models.QA, 4_500_000)
	

	prog1 := models.NewProgrammer("Marika", "Bayle", currentDate, 6_000_000, 500_000)
	prog2 := models.NewProgrammer("Red", "Hamilton", currentDate, 6_000_000, 500_000)
	prog3 := models.NewProgrammer("Lord", "Max", currentDate, 6_000_000, 500_000)
	prog4 := models.NewProgrammer("Blade", "Miquella", currentDate, 6_000_000, 500_000)

	sales1 := models.NewSales("Rex", "Lupus", currentDate, 3_000_000, 600_000)

	qa1 := models.NewQA("Rain", "Maker", currentDate, 4_500_000, 500_000)

	employees := []models.Employee{*&prog1.Employee, *&prog2.Employee, *&prog3.Employee, *&prog4.Employee, *&sales1.Employee}
	
	for _,v := range employees{
		fmt.Println(v.ToString())
		fmt.Println("")
	}

	fmt.Printf("Total salary all employee : %.2f\n", countSalary(employees))

	fmt.Printf("Profession of %s is %s\n\n", qa1.GetName(), qa1.GetProfession())

	progCount, salesCount, qaCount := countProfession(employees)
	fmt.Printf("Programmer Count : %d \nSales Count : %d \nQA count : %d", progCount, salesCount, qaCount)
}

func countSalary(person []models.Employee)float64{
	total := 0.0

	for _,v := range person{
		total += v.GetSalary()
	}

	return total
}

func countProfession(person []models.Employee)(int, int, int){
	progCount := 0
	salesCount := 0
	qaCount := 0

	for _,v := range person{
		if v.GetProfession() == "Programmer" {
			progCount += 1
		
		}else if v.GetProfession() == "Sales" {
			salesCount += 1
		
		}else{
			qaCount += 1
		}
	}

	return progCount, salesCount, qaCount
}