package models

import (
	"fmt"
	"math/rand"
	"time"
)

func NewEmployee(firstName string,
				lastName string,
				joinDate time.Time,
				role Roles,
				salary float64,
				) *Employee {

		return &Employee{
			empId: rand.Intn(101),
			firstName: firstName,
			lastName: lastName,
			joinDate: joinDate,
			role: role,
			salary: salary,
			}
}

func (emp *Employee) SetFirstName(firstName string) {
	emp.firstName = firstName
}

func (emp *Employee) SetLastName(lastName string) {
	emp.lastName = lastName
}

func (emp *Employee) SetJoinDate(joinDate time.Time) {
	emp.joinDate = joinDate
}

func (emp *Employee) SetRole(role Roles){
	emp.role = role
}

func (emp *Employee) SetSalary(salary float64) {
	emp.salary = salary
}

func (emp *Employee) ToString() string{
	return fmt.Sprintf("Employee : [%d, %s, %s, %s, %s, %.2f]",emp.empId, emp.firstName, emp.lastName, emp.joinDate, emp.role, emp.salary)
}


func (emp *Employee) GetSalary()float64{
	return emp.salary
}

func (emp *Employee) GetName()string{
	return fmt.Sprintf("%s %s", emp.firstName, emp.lastName)
}

// func (emp *Employee) GetProfession() string {
// 	return string(emp.role)
// }



// func (emp *Employee) SetTunjangan(tunjangan float64) {
// 	emp.tunjangan = tunjangan
// }

// func (emp *Employee) TotalSalary(salary float64, tunjangan float64) {
// 	emp.totalSalary = salary + tunjangan
// }

// func (emp *Employee) GetTunjangan()float64{
// 	return emp.tunjangan
// }