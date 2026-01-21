package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	arr := make([]int, 10)
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(arr []int, ch chan<- int) {
		defer wg.Done()
		defer close(ch)
		for i := range arr {
			arr[i] = rand.Intn(101)
			ch <- arr[i]
		}
	}(arr, ch)

	wg.Add(1)
	go func(ch <-chan int) {
		defer wg.Done()
		for i := range ch {
			fmt.Println(i * i)
		}
	}(ch)

	wg.Wait()
}
