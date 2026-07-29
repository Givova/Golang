package main

import (
	"fmt"
	"strings"
)

func WordCount(text string) map[string]int {
	 m := make(map[string]int)

	 words := strings.Fields(text)



	 for _, word := range words {
			m[word]++
	 } 

	 return m
}


func Unique(s []int) []int {

	was := make(map[int]bool)
	var res []int

	for _, n := range s {
		if !was[n]{
			was[n] = true
			res = append(res, n)
		}
	}

	return res

	// for i := 0; i < len(s); i++ {
	// 		if 
	// }
}

func Invert(m map[string]int) map[int]string  {
	
	mm := make(map[int]string)

	// s := make([]int, 10)

	// var str []string
	// var num []int

	for i, v := range m {
		mm[v] = i
	}

	return mm





}


func GroupBy(words []string, fn func(string) string) map[string][]string {

		result:= make(map[string][]string)

		for _, w := range words {
			key := fn(w)
			result[key] = append(result[key], w)
		}

		return result

}

func main() {
	// text := "go is great and go is fast"
	// fmt.Println(WordCount(text))
	// map[and:1 fast:1 go:2 great:1 is:2]


	// 2 задача
	// fmt.Println(Unique([]int{1, 2, 2, 3, 1, 4})) // [1 2 3 4]


	// 3 задача
	// m := map[string]int{"a": 1, "b": 2, "c": 3}
	// fmt.Println(Invert(m))
	// map[1:a 2:b 3:c]

	// 4 задача 

	words := []string{"go", "python", "c", "java", "js", "rust"}
	byLength := GroupBy(words, func(s string) string {
			if len(s) <= 2 {
					return "short"
			}
			return "long"
	})
	fmt.Println(byLength)
	// map[long:[python java rust] short:[go c js]]
}