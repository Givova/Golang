package main

import (
	"bufio"
	"sort"
	"strings"
)

// WordCount — пара слово + количество для сортировки.
type WordCount struct {
	Word  string
	Count int
}


// CountLines считает количество строк в тексте.
func CountLines(text string) int {
	// if text == "" {
	// 	return 0
	// }
	//lines := strings.Count(text, "\n" )
	// if text[len(text)-1] != '\n' {
	//       lines++
	//   }
	// return lines

	///////////////////////////////////
	// Способ через сканер !
	scanner := bufio.NewScanner(strings.NewReader(text))
	count := 0
	for scanner.Scan() {
		count++
	}
	return count

	// if text == ""{
	// 	return 0
	// }
	/////////////////////////////
	// способ через чистый подсчет символа \n
	// count := 1
	// for _, r := range text{
	// 	if r == '\n' {
	// 		count++
	// 	}
	// }
	// return count
}

// CountWords считает количество слов в тексте.
func CountWords(text string) int {
	if text == "" {
		return 0
	}

	str := strings.Fields(text)

	return len(str)
}

// CountChars считает количество символов (рун) в тексте.
func CountChars(text string) int {
	if text == "" {
		return 0
	}
	return len([]rune(text))

	// count := 0
	// for range text {
	// 	count++
	// }
	// return count
}

// WordFrequency возвращает карту слово→количество.
// Все слова приведены к нижнему регистру.
func WordFrequency(text string) map[string]int {
	m := make(map[string]int)

	str := strings.Fields(text)
	for _, v := range str {
		m[strings.ToLower(string(v))]++
	}

	return m
}

// TopWords возвращает топ-N слов по частоте.
// Результат — слайс пар (слово, количество), отсортированный по убыванию.
func TopWords(freq map[string]int, n int) []WordCount {

	var slice []WordCount

	for key, value := range freq {
		el := WordCount{key, value}
		slice = append(slice, el)

	}

	sort.Slice(slice, func(i, j int) bool {
		if slice[i].Count == slice[j].Count {
    return slice[i].Word < slice[j].Word // алфавит
}
		return slice[i].Count > slice[j].Count
	})


	if n >= len(slice){
		return slice
	}
	return slice[0:n]
}
