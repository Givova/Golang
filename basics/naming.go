package basics

import (
	"strings"
	"unicode"
)

// type EmployeeGoogle struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// }

// type EmployeeApple struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// }

// func concatPlus(strs []string) string {
// 	result := ""
// 	for _, s := range strs {
// 		result += s // каждое сложение создает новую строку

// 	}
// 	return result
// }

// func main() {
// 	//PascalCase
// 	//Eg. CalculateArea, UserInfo, NewHTTPRequest
// 	// Struct, interface, enums

// 	//snake_case
// 	// user_id, first_name

// 	//UPPERCASE
// 	// Use case in constant

// 	//mixedCase
// 	// Eg. javaScript, htmlDocument, isValid

// 	// const MAXRETRIES = 5

// 	// var employeeID = 1001
// 	// fmt.Println("EmployeeID: ", employeeID)

// }

// func concatBuilder(strs []string) string {
// 	var builder strings.Builder

// 	// Оптимизация, предварительно выделяем память
// 	totalLength := 0
// 	for _, s := range strs {
// 		totalLength += len(s)
// 	}
// 	builder.Grow(totalLength)  //Важная оптимизация

// 	//конкатенация
// 	for _, s := range strs {
// 		builder.WriteString(s)
// 	}

// 	return builder.String()

// }

func reverseString(s string) string {


	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left < right{
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return string(runes)
}



func isPalindrome(s string) bool {

	var cleaned [] rune

	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, r)
		}
	}

	left, right := 0, len(cleaned)-1
	for left < right {
		if cleaned[left] != cleaned[right]{
				return false		
		}
		left++
		right--
	}
	return true
}

func firstUniqChar(s string) int {
	//Создаем мапу для подсчета частот
	freq := make(map[rune]int)

	//Преобразуем строку в слайс рун для корректного подсчета
	runes := []rune(s)

	//Певрый проход: подсчет частот
	for _, char := range runes {
		freq[char]++
	}

	//Второй проход: поиск первого уникального
	for i, char := range runes {
		if freq[char] == 1 {
			return i //Теперь i - это корректный индекс символа
		}	
	}

	return -1

}


