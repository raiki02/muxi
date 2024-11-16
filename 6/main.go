// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// var flag = true

// var wg sync.WaitGroup

// func get(ch chan int) {
// 	num := 0
// 	ok := true
// 	for ok {
// 		_, flag = <-ch
// 		num++
// 		if flag == false {
// 			break
// 		}
// 	}
// 	fmt.Println("Moved ", num)

// }

// func main() {
// 	runtime.GOMAXPROCS(10)
// 	cap := make(chan int, 2000)
// 	//1 -> 2000 pool
// 	go func() {
// 		for i := 0; i < 2000; i++ {
// 			cap <- 1
// 		}
// 	}()

// 	//car
// 	cars := make(chan int, 10)

// 	wg.Wait()
// }

package main

func get(ch chan int) {
	flag := true

}
func main() {

	ch := make(chan int, 2000)
	for i := 0; i < 2000; i++ {
		ch <- 1
	}

	go func() {
		for i := 0; i < 10; i++ {
			get(ch)
		}
	}()
}
