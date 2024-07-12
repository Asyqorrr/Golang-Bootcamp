package models

import (
	"time"
)

type SalesInterface interface{
	SetTunjanganSales()
	GetTotalIncome()
}

type Saless struct {
	Employee

	tunjanganSales float64
}

func NewSales(
	firstName string,
	lastName string,
	joinDate time.Time,
	salary float64,
	tunjanganSales float64) *Saless{

		return &Saless{
			Employee:*NewEmployee(firstName,
									lastName,
									joinDate,
									Sales,
									salary), 
			tunjanganSales:tunjanganSales}
}

func (sales *Saless) SetTunjanganSales(tunjanganSales float64){
	sales.tunjanganSales = tunjanganSales
}

func (sales *Saless) GetTotalIncome()float64{
	return sales.salary + sales.tunjanganSales
}
