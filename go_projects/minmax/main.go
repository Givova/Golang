package main

import "fmt"

func clamp(val *int, min, max int)  {

	if *val < min {
		  *val = min
	} else if  *val > max {
		*val = max
	}

	}

func main() {
	x := 150
	clamp(&x, 0, 100)
	fmt.Println(x) // 100

	y := -5
	clamp(&y, 0, 100)
	fmt.Println(y) // 0

	z := 50
	clamp(&z, 0, 100)
	fmt.Println(z) // 50
}