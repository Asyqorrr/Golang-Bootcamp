package models

import "fmt"

type Boat struct {
	Vehicle

	OrderPerHours float64
}

func NewBoat(
	noPolice string,
	year string,
	price float64,
	tax float64,
	seat string,
	OrderPerHours float64) *Boat {

	return &Boat{
		Vehicle: *NewVehicle(
			noPolice,
			VBoat,
			year,
			price,
			tax,
			seat),
		OrderPerHours: OrderPerHours,
	}
}

type BoatInterface interface {
	SetOrderPerHours(OrderPerHours float64)
}

func (boat *Boat) SetOrderPerHours(OrderPerHours float64) {
	boat.OrderPerHours = OrderPerHours
}

func (boat *Boat) Info() string {
	return fmt.Sprintf("\nNo Police : %s\nVehicle Type : %s\nPrice : %.1f\nTax(perYear) :%.1f\nSeat :%s\nOrderPerHours Income :%.2f\n", boat.noPolice, boat.vehType, boat.price, boat.tax, boat.seat, boat.OrderPerHours)
}

func (boat *Boat) GetTotalIncome() float64 {
	return boat.OrderPerHours
}