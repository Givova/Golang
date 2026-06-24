package main

import (
	"fmt"
)

func equalFloat(a, b float64) bool {
	diff := a - b

	// берём модуль разности
	if diff < 0 {
		diff = -diff
	}

	// сравниваем с допустимой погрешностью
	return diff < 0.000000001
}

func main() {
    a := 0.1
    b := 0.2
    c := 0.3

    fmt.Println("Функция equalFloat сравнивает корректно:")
    fmt.Println("equalFloat(0.1+0.2, 0.3):", equalFloat(a+b, c)) // true
    
    // Обычные числа тоже работают
    fmt.Println("equalFloat(1.5, 1.5):", equalFloat(1.5, 1.5))     // true
    fmt.Println("equalFloat(1.5, 1.5000001):", equalFloat(1.5, 1.5000001)) // false

    fmt.Println("Прямое сравнение некорректно:")
    fmt.Println("0.1 + 0.2 == 0.3:", a+b == c)        // false
    fmt.Println("0.1 + 0.2 =", a+b)                   // 0.30000000000000004
}