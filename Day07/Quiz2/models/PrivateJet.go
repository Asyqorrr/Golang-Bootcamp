package models

import "fmt"

type PrivateJet struct {
	Vehicle

	OrderPerHours float64
}

func NewPrivateJet(
	noPolice string,
	year string,
	price float64,
	tax float64,
	seat string,
	OrderPerHours float64) *PrivateJet {
	return &PrivateJet{
		Vehicle: *NewVehicle(
			noPolice,
			VPrivateJet,
			year,
			price,
			tax,
			seat),
		OrderPerHours: OrderPerHours,
	}
}

type PrivateJetInterface interface {
	SetOrderPerHours(OrderPerHours float64)
	GetTotalIncome() float64
}

func (jet *PrivateJet) SetOrderPerHours(OrderPerHours float64) {
	jet.OrderPerHours = OrderPerHours
}

func (jet *PrivateJet) GetTotalIncome() float64 {
	return jet.OrderPerHours
}

func (jet *PrivateJet) Info() string {
	return fmt.Sprintf("\nNo Police : %s\nVehicle Type : %s\nPrice : %.1f\nTax(perYear) :%.1f\nSeat :%s\nOrderPerHours Income :%.2f\n", jet.noPolice, jet.vehType, jet.price, jet.tax, jet.seat, jet.OrderPerHours)
}
