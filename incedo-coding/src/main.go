package main

import (
	"fmt"
	"sync"
)

func main() {
	even := 0
	odd := 1
	limit := 10
	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		count := 0
		for count < limit {
			even = even + 2
			ch <- even
			count++
		}
	}()

	go func() {
		defer wg.Done()
		count := 0
		for count < limit {
			odd = odd + 2
			ch <- odd
			count++
		}
	}()

	go func() {
		defer wg.Done()
		count := 0
		for count < (limit * 2) {
			v := <-ch
			fmt.Println(v)
		}
	}()

	wg.Wait()
}

func fib(n int) {

	arr := retfib(n, 0, 1)
	for _, v := range arr {
		fmt.Println(v)
	}
}

func retfib(n, first, next int) []int {
	retArr := make([]int, 0)
	i := 0
	for i < n {
		retArr = append(retArr, first)
		first, next = next, (first + next)
		i++
	}
	return retArr
}
