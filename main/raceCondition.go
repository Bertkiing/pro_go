package main

import (
	"sync"
	"runtime"
	"fmt"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	fmt.Println("Final Counter :", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count ++ {
		value := counter

		/**
			Gosched()函数，用于将goroutine从当前线程退出...给其它goroutine运行的机会。
		 */
		runtime.Gosched()

		value ++

		counter = value
	}
}
