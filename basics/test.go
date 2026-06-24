package basics

import (
	"fmt"
)

func main() {
	x:=2
	if x % 2 == 0 {
		fmt.Println( x, "- четное число")
	} else {
		fmt.Println( x, "- нечетное число")
	}
}