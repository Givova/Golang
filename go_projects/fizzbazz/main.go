package main

import "fmt"

var n int

func main() {

	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	if n == 0 || n < 0{
		fmt.Print("Дятел, введи положительное число!")
	}

	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBizz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Bizz")

		default:
			fmt.Println(i)
		}
	}
}
