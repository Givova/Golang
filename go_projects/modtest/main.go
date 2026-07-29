package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)
 

func main(){
caser := cases.Title(language.Russian)
fmt.Println(caser.String("привет мир")) // Привет Мир
}
