package models

import (
	"fmt"
	"math/rand"
)

type Variant string

const (
	VSuv        Variant = "SUV"
	VTaxi       Variant = "Taxi"
	VPrivateJet Variant = "Private Jet"
	VBoat  		Variant = "Boat"
)

type Vehicle struct {
	vehId    int
	noPolice string
	vehType  Variant
	year     string
	price    float64
	tax      float64
	seat     string
}

//generate interface:
// emp *Employee EmployeeInterface , tekan enter

type VehicleInterface interface{
	SetNoPolice(noPolice string)
	SetVehType(vehType Variant)
	SetYear(year string)
	SetPrice(price float64)
	SetTax(tax float64)
	SetSeat(seat string)
	GetTax() float64
	GetTotalIncome() float64
	GetType() string
	Info() string
}

func NewVehicle(noPolice string, vehType Variant, year string, price float64, tax float64, seat string) *Vehicle {

	return &Vehicle{
		vehId: rand.Intn(100),
		noPolice: noPolice,
		vehType: vehType,
		year: year,
		price: price,
		tax: tax,
		seat: seat}
}

func (veh *Vehicle) SetNoPolice(noPolice string) {
	veh.noPolice = noPolice
}

func (veh *Vehicle) SetVehType(vehType Variant) {
	veh.vehType = vehType
}

func (veh *Vehicle) SetYear(year string) {
	veh.year = year
}

func (veh *Vehicle) SetPrice(price float64) {
	veh.price = price
}

func (veh *Vehicle) SetTax(tax float64) {
	veh.tax = tax
}

func (veh *Vehicle) SetSeat(seat string) {
	veh.seat = seat
}

func (veh *Vehicle) Info()string {
	return fmt.Sprintf("\nNo Police : %s\nVehicle Type : %s\nPrice : %.1f\nTax(perYear) :%.1f\nSeat :%s", veh.noPolice, veh.vehType, veh.price, veh.tax, veh.seat)
}

func (veh *Vehicle) GetTax() float64{
	return veh.tax
}

func (veh *Vehicle) GetTotalIncome() float64{
	return 0.00
}

func (veh *Vehicle) GetType() string{
	return string(veh.vehType)
}

