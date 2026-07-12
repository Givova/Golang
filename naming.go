package main

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}


type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}


func main() {
	//PascalCase
	//Eg. CalculateArea, UserInfo, NewHTTPRequest
	// Struct, interface, enums

	//snake_case
	// user_id, first_name

	//UPPERCASE
	// Use case in constant

	//mixedCase
	// Eg. javaScript, htmlDocument, isValid

	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)
}
