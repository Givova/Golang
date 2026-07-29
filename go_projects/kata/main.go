package main

import (
	"fmt"
	"strings"
)


func SpinWords(str string) string {
  words := strings.Fields(str)  //["sds" 'sds' "sda"]
  
  for i, word := range words{ // проходимся по всем словам в слайсе
    if len(word) >= 5 {       // если больше 5 символов
			runes:= []rune(word)    // создаем слайс рун

      for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1{
        runes[i], runes[j] = runes[j], runes[i]              // меняем местами буквы по задаче 
        
      }
			 words[i] = string(runes)         //Перевёрнутые руны снова собираем в строку и кладём обратно в слайс на место старого слова.
    }
  }
	
return strings.Join(words, " ")   // собираем через пробел в одну строку
}





// func OneSpinWords(str string) string {
//   s := strings.Split(str, " ")
  
//   for i, v := range s {
//     if len(v) >= 5 {
//       res := ""
//       for _, r := range v {
//         res = string(r) + res
//       }
//       s[i] = res
//     }
//   } 
  
//   return strings.Join(s, " ")
// }

func main() {
	fmt.Println(SpinWords("Hey fellow Warriors"))
}