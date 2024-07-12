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
		GenerateEmployeesWithChannel(200)
	}
}