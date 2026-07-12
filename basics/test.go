package basics

import (
	"fmt"
)

func bas() {
	x:=2
	if x % 2 == 0 {
		fmt.Println( x, "- четное число")
	} else {
		fmt.Println( x, "- нечетное число")
	}
}