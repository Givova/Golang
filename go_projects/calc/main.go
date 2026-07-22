package main

import "fmt"

// var a float64
// var b float64
// var op string

func add(a, b float64) float64 {
	return a + b
}

func substract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false
	}

	return a / b, true
}

func readNumber(prompt string) float64{

	var n float64
  fmt.Print(prompt)
  fmt.Scan(&n)
  return n
}




func main() {
	

	a := readNumber("Первое число: ")
	b := readNumber("Второе число: ")

	var op string
	fmt.Print("Операция(+, -, *, /): ")
	fmt.Scan(&op)

	var result float64
	var ok bool

	switch op {
	case "+":
		result = add(a, b)
		ok = true
	

	case "-":
		result = substract(a, b)
		ok = true
	

	case "*":
		result = multiply(a, b)
		ok = true
	
	case "/":
		result, ok = divide(a, b)

	default: 
	 fmt.Println("Неизвестная операция!")
	 return
	}


	if !ok {
		fmt.Println("Ошибка - Деление на ноль")
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", a, op, b, result)

}
