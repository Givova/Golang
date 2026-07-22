package main

import (
	"fmt"
	"sync"
	"time"
)

// const pi = 3.14
// const GRAVITY = 9.81

// func main() {
// 	const days int = 7

// 	const (
// 		monday   = 1
// 		tuesday  = 2
// 		wenesday = 3
// 	)

// }

// func climbStairs(n int) int {
//     if n == 1 {
//         return 1
//     }
//     if n == 2 {
//         return 2
//     }

//     prev2 := 1 // dp[i-2]
//     prev1 := 2 // dp[i-1]

//     for i := 3; i <= n; i++ {
//         cur := prev1 + prev2
//         prev2 = prev1
//         prev1 = cur
//     }
//     return prev1
// }

// func doWork(ctx context.Context) {
// 	newCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
// 	defer cancel()

// 	log.Println("starting working ... ")

// 	for {
// 		select {
// 		case <-newCtx.Done():
// 			log.Printf("ctx done: %v", ctx.Err())
// 			return
// 		default:
// 			log.Println("working ... ")
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	doWork(ctx)

// 	// fmt.Println(" Ступенек необходимо: ")
// 	// fmt.Println(climbStairs(6))
// 	// sum := 0

// 	// for {
// 	// 	sum+= 10
// 	// 	fmt.Println("Sum:", sum)
// 	// 	if sum >= 50{
// 	// 		break
// 	// 	}
// 	// }

// 	// source:= rand.NewSource(time.Now().UnixNano())
// 	// random := rand.New(source)

// 	// target := random.Intn(100) +1

// 	// fmt.Println("Welcome to the Guessing Game")
// 	// fmt.Println("I have chosen a number between 1 and 100")
// 	// fmt.Println("Can you guess what it is?")

// 	// var guess int
// 	// for{
// 	// 	fmt.Println("Enter your guess: ")
// 	// 	fmt.Scanln(&guess)

// 	// 	//Check it the guess in correct

// 	// 	if guess == target{
// 	// 		fmt.Println("Congratulations! You guessed the correct number!")
// 	// 		break
// 	// 	} else if guess  < target {
// 	// 		fmt.Println("Too low! Try guessing a higher number.")
// 	// 	} else if guess > target {
// 	// 		fmt.Println("Too high! Try guessing a lower number.")
// 	// 	}

// 	// }

// 	// switch expression {
// 	// case condition:

// 	// }
// }

type SafeCounter struct {
	mu sync.Mutex
	v map[string]int 
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++ 
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c:= SafeCounter{
		v: make(map[string]int),
	}

	for i:=0; i<1000; i++{
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)

	fmt.Println(c.Value("somekey"))
}

// func main() {
//     c := SafeCounter{v: make(map[string]int)}
//     for i := 0; i < 1000; i++ {
//         go c.Inc("somekey")
//     }

// 	time.Sleep(time.Second)
// 	fmt.Println(c.Value("somekey"))
// }