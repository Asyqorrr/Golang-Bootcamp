package models

import (
	"time"
)

type QAInterface interface{
	SetTunjanganQA()
	GetTotalIncome()
	GetProfession()
}

type QAs struct {
	Employee

	tunjanganQA float64
}

func NewQA(
	firstName string,
	lastName string,
	joinDate time.Time,
	salary float64,
	tunjanganQA float64) *QAs{

		return &QAs{
			Employee:*NewEmployee(firstName,
									lastName,
									joinDate,
									QA,
									salary), 
			tunjanganQA: tunjanganQA}
}

func (qa *QAs) SetTunjanganQA(tunjanganQA float64){
	qa.tunjanganQA = tunjanganQA
}

func (qa *QAs) GetTotalIncome()float64{
	return qa.salary + qa.tunjanganQA
}

func (qa *QAs) GetProfession()string {
	return "QA"
}