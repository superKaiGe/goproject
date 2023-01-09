package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func main() {
	dilbert.Salary = 5000
	position := &dilbert.Position
	*position = "Senior" + *position
	var employeeOfTheMonth *Employee = &dilbert

	employeeOfTheMonth.Position += " (proactive team player)"
	//后面一条语句等价于
	(*employeeOfTheMonth).Position += " (proactive team player)"
	fmt.Println(EmployeeById(dilbert.ManagerID).Position)
}
func EmployeeById(id int) *Employee {
	return &Employee{}
}
