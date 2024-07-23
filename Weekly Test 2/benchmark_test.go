package main

import "testing"

func BenchmarkAllTwoGenerator(b *testing.B) {
	b.Run("BenchGenerateEmployees", BenchGenerateEmployees)
	b.Run("BenchGenerateEmployeesWithChannel", BenchGenerateEmployeesWithChannel)
}

func BenchGenerateEmployees(b *testing.B){
	for i := 0; i < b.N; i++{
	GenerateEmployees(200)
	}
}

func BenchGenerateEmployeesWithChannel(b *testing.B){
	for i := 0; i < b.N; i++{
		go GenerateEmployeesWithChannel(200)
	}
}

func BenchmarkEmployeeTotalSalariesWithChannel(b *testing.B){
	employees := GenerateEmployeesWithChannel(5000)

	for i := 0; i < b.N; i++{
		go EmployeeTotalSalariesWithChannel(employees)
	}
}

func BenchmarkEmployeeTotalSalaries(b *testing.B){
	employees := GenerateEmployeesWithChannel(5000)

	for i := 0; i < b.N; i++{
		EmployeeTotalSalaries(employees)
	}
}

//1.236.922.057 without channel
//1.167.361.755 with channel