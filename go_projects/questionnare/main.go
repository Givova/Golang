package main

import (
	"fmt"
	"strconv"
)

// var name string
// var age int
// var high float64

func main() {

	// fmt.Println("Введите имя: ")
	// fmt.Scan(&name)
	// fmt.Println("Введите возраст: ")
	// fmt.Scan(&age)
	// fmt.Println("Введите рост: ")
	// fmt.Scan(&high)

	// fmt.Printf("Имя: %s, Возраст: %d, Рост: %.1f", name, age, high)

	// x := 5
	// y := 2.5
	// fmt.Println(float64(x) + float64(y))

	num := 42
	text := strconv.Itoa(num)
	fmt.Println(text) // хотим "42", а получаем...

}
