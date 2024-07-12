package models

import "fmt"

type Taxi struct {
	Vehicle

	order      int
	orderPerKM float64
}

func NewTaxi(noPolice string,
	year string,
	price float64,
	tax float64,
	seat string,
	order int,
	orderPerKM float64) *Taxi {

	return &Taxi{
		Vehicle: *NewVehicle(
			noPolice,
			VTaxi,
			year,
			price,
			tax,
			seat,
		),
		order:      order,
		orderPerKM: orderPerKM}
}

type TaxiInterface interface {
	SetOrder(order int)
	SetOrderPerKM(orderPerKM float64)
}

func (taxi *Taxi) SetOrder(order int) {
	taxi.order = order
}

func (taxi *Taxi) SetOrderPerKM(orderPerKM float64) {
	taxi.orderPerKM = orderPerKM
}

func (taxi *Taxi) GetTotalIncome() float64 {
	return float64(taxi.order) * taxi.orderPerKM
}

func (taxi *Taxi) Info() string {
	return fmt.Sprintf("\nNo Police : %s\nVehicle Type : %s\nPrice : %.1f\nTax(perYear) :%.1f\nSeat :%s\nOrder :%.d\nOrderPerKM Income:%.2f", taxi.noPolice, taxi.vehType, taxi.price, taxi.tax, taxi.seat, taxi.order, taxi.orderPerKM)
}