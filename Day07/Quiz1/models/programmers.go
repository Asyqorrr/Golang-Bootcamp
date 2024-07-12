package models

import (
	"time"
)

type Programmers struct {
	Employee

	tunjanganProg float64
}

type ProgrammersInterface interface{
	SetTunjanganProg()
	GetTotalIncome()
}

func NewProgrammer(
	firstName string,
	lastName string,
	joinDate time.Time,
	salary float64,
	tunjanganProg float64) *Programmers {

	return &Programmers{
		Employee:*NewEmployee(firstName,
								lastName,
								joinDate,
								Programmer,
								salary),
		tunjanganProg: tunjanganProg}
}

func (programmer *Programmers) SetTunjanganProg(tunjanganProg float64){
	programmer.tunjanganProg = tunjanganProg
}

func (programmer *Programmers) GetTotalIncome()float64{
	return programmer.salary + programmer.tunjanganProg
}

func (programmer *Programmers) GetProfession() string{
	return "Programmer"
}

// func (programmer *Programmers) Info() string {
// 	return fmt.Sprintf("%v, ")
// }
