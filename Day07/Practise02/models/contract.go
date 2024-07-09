package models

import (
	"fmt"
	"time"
)

//struct contract
type Contract struct {
	Employee //embedded
	overTime float64
}

// interface Contract
type ContractInterface interface {
	SetOvertime()
}

// constructor
func NewContract(fullName string,
	dateOfBirth time.Time,
	salary float64,
	overtime float64) *Contract {

	return &Contract{*NewEmployee(fullName, dateOfBirth, salary, Contracts), overtime}
}

//set function
func (contract *Contract) SetOvertime(overtime float64) {
	contract.overTime = overtime
}

//get function
func (contract *Contract) Info() string {
	return fmt.Sprintf("%v, Overtime : [%.2f]", contract.Employee, contract.overTime)
}
