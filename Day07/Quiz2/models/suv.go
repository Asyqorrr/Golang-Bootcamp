package models

import "fmt"

type SUV struct {
	Vehicle

	rentIncome   float64
	driverIncome float64
}

func NewSUV(
	noPolice string,
	year string,
	price float64,
	tax float64,
	seat string,
	rentIncome float64,
	driverIncome float64) *SUV {

	return &SUV{
		Vehicle: *NewVehicle(noPolice,
			VSuv,
			year,
			price,
			tax,
			seat,
		),
		rentIncome:   rentIncome,
		driverIncome: driverIncome}
}

type SUVInterface interface {
	SetRentIncome(rentIncome float64) 
	SetDriverIncome(driverIncome float64) 
	GetType() string
	GetTotalIncome() float64
}

func (suv *SUV) SetRentIncome(rentIncome float64) {
	suv.rentIncome = rentIncome
}

func (suv *SUV) SetDriverIncome(driverIncome float64) {
	suv.driverIncome = driverIncome
}

func (suv *SUV) Info() string {
	return fmt.Sprintf("\nNo Police : %s\nVehicle Type : %s\nPrice : %.1f\nTax(perYear) :%.1f\nSeat :%s\nRent Income :%.2f\nDriver Income:%.2f\n------------------------------------", 
	suv.noPolice, suv.vehType, suv.price, suv.tax, suv.seat, suv.rentIncome, suv.driverIncome)
}

func (suv *SUV) GetTotalIncome()float64{
	return suv.rentIncome + suv.driverIncome
}

func (suv *SUV) GetType() string{
	return string(suv.vehType)
}

