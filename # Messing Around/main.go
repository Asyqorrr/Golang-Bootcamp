package main

import "fmt"

// Define a base struct
type Address struct {
    Street string
    City   string
}

// Constructor function for Address
func NewAddress(street, city string) *Address {
    return &Address{
        Street: street,
        City:   city,
    }
}

// Define another base struct
type Person struct {
    Name string
    Age  int
}

// Constructor function for Person
func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}

// Define a struct that embeds Address and Person
type Employee struct {
    Person
    Address
    position string
    salary   float64
}

// Constructor function for Employee
func NewEmployee(name string, age int, street, city, position string, salary float64) *Employee {
    return &Employee{
        *NewPerson(name, age),
        *NewAddress(street, city),
        position,
        salary,
    }
}

func main() {
    // Use the constructor to create a new instance of Employee
    // name := []string {"Halimawan", "Merkava", "Marika", "Radahn", "Daenarys", "Stark", "Uzumaki"}
    
    a := 1
    b := 0

    if a < 3 {
        b = 3
    }

    fmt.Println(b)
}