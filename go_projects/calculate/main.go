package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// = - * /
func add(a, b float64) float64 {
  return a + b
}

func subtract(a, b float64) float64 {
  return a - b
}
func multiply(a, b float64) float64 {
  return a * b
}
func divide(a, b float64) (float64, error) {
  if b == 0 {
    return 0, errors.New("Деление на ноль")
  }

  return a / b, nil
}

func parseExpression(input string) (float64, string, float64, error) {
  //str   2 + 3 


	var a, b float64
	var op string


 
	n, err := fmt.Sscanf(input, "%f %s %f", &a, &op, &b)	
	if err != nil || n != 3 {
		return 0, "", 0, errors.New("неверный формат")
	}

  // strv := strings.Fields(input)
  // n, err := strconv.ParseFloat(strv[ ], 64)
  // input, err := fmt.Sscanf("%s", &input)
  // if err != nil {
  //   return 0, "", 0, errors.New("Ошибка")
  // }

  return a, op, b, nil
}


func calculate(a float64, op string, b float64) (float64, error){



	switch op{
	case "+":
		// fmt.Printf(" %.2f %v %.2f = %v ", a, op, b,add(a,b))
		return add(a,b), nil
	case "-":
		// fmt.Printf(" %.2f %v %.2f = %v ", a, op, b, subtract(a,b))
		return subtract(a,b), nil
	case "*":
		// fmt.Printf(" %.2f %v %.2f = %v ", a, op, b, multiply(a,b))
		return multiply(a,b), nil
	case "/":	
		res, err := divide(a,b)
		if err != nil{
			return 0, err
		}
		// fmt.Printf(" %.2f %v %.2f = %v ", a, op, b, res)
		return res, nil
	default:
        return 0, fmt.Errorf("неизвестная операция: %q", op)
  
	}
}

func main() {
  defer fmt.Println("Пока!")
 
    fmt.Println("Калькулятор. Введи выражение (или \"exit\" для выхода).")
    fmt.Println()


	scanner := bufio.NewScanner(os.Stdin)

	for {
		 
		if !scanner.Scan(){ //«Если не удалось прочитать строку — выйти из цикла»
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if strings.EqualFold(line, "exit"){
			return
		}
		if line == ""{
			continue
		}

		a, op, b, err := parseExpression(line) 

		if err != nil{
			fmt.Printf("Ошибка: %v\n\n", err)
			continue
		} 


		result, err := calculate(a, op, b)
		if err != nil{
			fmt.Printf("Ошибка: %v\n\n", err)
      continue
		}

		fmt.Printf("Результат: %.2f\n\n", result)
	}


	// if err := scanner.Err();
	//  if err != nil {
	// 	fmt.Println("Ошибка чтения:", err)
	// }

	if scanner.Err() != nil {                               // эта запись короче!
    fmt.Println("Ошибка чтения:", scanner.Err())
	}

	
}
