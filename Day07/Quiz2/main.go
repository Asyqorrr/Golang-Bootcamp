package main

import (
	"day07/quiz2/models"
	"fmt"
)


func main() {

	//  ---------------------------------------------------------------------------------------------
	suv1 := models.NewSUV("D 1001UM ", "2010", 350_000_000, 3_500_000, "4", 500_000, 150_000)
	suv2 := models.NewSUV("D 1002UM ", "2010", 350_000_000, 3_500_000, "4", 500_000, 150_000)
	suv3 := models.NewSUV("D 1003UM ", "2015", 350_000_000, 3_500_000, "5", 500_000, 150_000)
	suv4 := models.NewSUV("D 1004UM ", "2015", 350_000_000, 3_500_000, "5", 500_000, 150_000)

	taxi1 := models.NewTaxi("D 11 UK", "2002", 175_000_000, 1_750_000, "4", 45, 4_500)
	taxi2 := models.NewTaxi("D 12 UK", "2015", 275_000_000, 2_250_000, "4", 75, 4_500)
	taxi3 := models.NewTaxi("D 13 UK", "2020", 275_000_000, 2_750_000, "4", 90, 4_500)

	plane1 := models.NewPrivateJet("ID8089", "2015", 150_000_000_000, 1_500_000_000, "12", 55_000_000)
	plane2 := models.NewPrivateJet("ID8099", "2018", 175_000_000_000, 1_750_000_000, "15", 75_000_000)

	boat1 := models.NewBoat("Boat18", "2020", 2_000_000_000, 20_000_000, "12", 10_000_000)
	//  ---------------------------------------------------------------------------------------------

	vehicles := []models.VehicleInterface{suv1,suv2,suv3,suv4,taxi1, taxi2,taxi3,plane1,plane2,boat1}
	
	suvIncome, taxiIncome, jetIncome, boatIncome := TotalIncome(vehicles)
	fmt.Printf("\nSUV income: %.2f",suvIncome)
	fmt.Printf("\nTaxi income: %.2f",taxiIncome)
	fmt.Printf("\nJet income: %.2f",jetIncome)
	fmt.Printf("\nBoat income: %.2f\n---------------------",boatIncome)

	totalTax := totalTax(vehicles)
	fmt.Printf("\nTotalTax = %.2f\n---------------------", totalTax)
	
	totalCar, totalPlane, totalBoat := countVehicle(vehicles)
	fmt.Printf("\nTotal Car : %d", totalCar)
	fmt.Printf("\nTotal Plane : %d", totalPlane)
	fmt.Printf("\nTotal Boat : %d", totalBoat)
}

//so we create an interfaces and declare what methods we want to use in that interfaces?
//and after that i create a slices of that particular interfaces and i can kind of 
// iterates different structs on it right?

func TotalIncome(vehicles []models.VehicleInterface)(float64, float64, float64, float64){
	suvIncome := 0.0
	taxiIncome := 0.0
	jetIncome := 0.0
	boatIncome := 0.0

	for _,v := range vehicles{
		if v.GetType() == "SUV"{
			suvIncome += v.GetTotalIncome()
		
		}else if v.GetType() == "Taxi"{
			taxiIncome += v.GetTotalIncome()
		
		}else if v.GetType() == "Private Jet"{
			jetIncome += v.GetTotalIncome()
		
		}else {
			boatIncome += v.GetTotalIncome()
		}
	}
	return suvIncome, taxiIncome, jetIncome, boatIncome
}

func totalTax(vehicles []models.VehicleInterface)float64{
	totalTax:= 0.0

	for _,v := range vehicles {
		totalTax += v.GetTax()
	}
	return totalTax
}

func countVehicle (vehicles []models.VehicleInterface)(int, int, int){
	totalCar :=0
	totalPlane :=0
	totalBoat :=0

	for _,v := range vehicles{
		if v.GetType() == "SUV" || v.GetType() == "Taxi"{
			totalCar++
		
		}else if v.GetType() == "Private Jet"{
			totalPlane++

		}else {
			totalBoat++
		}
	}
	return totalCar, totalPlane, totalBoat
}
