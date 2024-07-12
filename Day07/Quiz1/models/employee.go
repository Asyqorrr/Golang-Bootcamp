package models

import "time"

type (
	Roles string
)

const (
	Programmer Roles 	= "Programmer"
	Sales Roles			= "Sales"
	QA 	Roles			= "QA"
)

type Employee struct {
	empId     	int
	firstName 	string
	lastName  	string
	joinDate  	time.Time
	role		Roles
	salary 		float64
}

type EmployeeInterface interface {
	SetFirstName(firstName string)
	SetLastName(lastName string)
	SetJoinDate(joinDate time.Time)
	SetRole(role Roles)
	SetSalary(salary float64)
	ToString()
}

type Info interface {
	ToString()
}