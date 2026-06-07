package main

import "fmt"

var  middleName = "Farhat"    
// только внутри функции можно объявить переменную


func main() {
// 	var age int
// 	var name string = `John`
// 	var name1 = "jan"

// 	count := 10
// 	lastName := "Doe"
middleName := "John"



	// Default values
	//Numeric types: 0
	//String types: ""
	//Boolean types: false
	//Pointer types: nil
	//Slice types: nil
	//Map types: nil
	//Channel types: nil
	//Function types: nil
	//Interface types: nil
	//Struct types: nil


	fmt.Println(middleName)
}

func printname() {
	firstName := "Michael"
	fmt.Println(firstName)
}