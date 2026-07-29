package main

import "fmt"

func Filter(s []int, fn func(int) bool) []int {
	var ans []int

	for i, n := range s {
		res := fn(n)
		if res {
			ans = append(ans, s[i])
		}
	}

	return ans
}

// func Filter(s []int, fn func(int) bool) []int {
//     var result []int

//     for _, v := range s {
//         if fn(v) {
//             result = append(result, v)
//         }
//     }
//     return result
// }

func Chunk(s []int, size int) [][]int {

	var result [][]int

	for i := 0; i < len(s); i += size {
		end := i + size
		if end > len(s) {
			end = len(s)
		}
		result = append(result, s[i:end])
	}
	return result
}

func main() {
	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	// fmt.Println(evens) // [2 4 6 8]

	// big := Filter(nums, func(n int) bool { return n > 5 })
	// fmt.Println(big) // [6 7 8]

	fmt.Println(Chunk([]int{1, 2, 3, 4, 5}, 2)) // [[1 2] [3 4] [5]]
	fmt.Println(Chunk([]int{1, 2, 3}, 5))       // [[1 2 3]]

}