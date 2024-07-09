package models

import (
	"fmt"
	"time"
)

type Permanent struct {
	Employee //embedded
	//emps      Employee //not emmbedded
	insurance float64
}

// interface Permanent
type PermanentInterface interface {
	setInsurance()
}

// constructor
func NewPermanent(fullName string,
	dateOfBirth time.Time,
	salary float64,
	insurance float64) *Permanent {

	return &Permanent{*NewEmployee(fullName, dateOfBirth, salary, Permanents), insurance}
}

//set function
func (permanent *Permanent) SetInsurance(insurance float64){
	permanent.insurance = insurance
}

//get function
func (permanent *Permanent) Info() string {
	return fmt.Sprint("%v, Insurance : [%.2f]", permanent.Employee, permanent.insurance)
}