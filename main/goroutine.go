package main

import (
	"fmt"
	"sync"
	"runtime"
)

/**
	runtime包提供了修改Go语言运行时配置参数的能力。
	函数NumCPU()返回可以使用的物理处理器的数量。
	显然，只有在有多个逻辑处理器且可以同时让每个goroutine运行在一个可用的物理处理器上的时候，goroutine才会真正并行运行。
 */



/**
	基于调度器的内部算法，一个正运行的goroutine在工作结束前，可以被停止并重新调度。
	目的：防止某个goroutine长时间占用逻辑处理器。当goroutine占用时间过长时，调度器会停止当前正运行的goroutine，
并给其它可运行的goroutine运行的机会。
 */

func main() {
	/**
		Go的标准库的runtime包里的GOMAXPROCS()：指定调度器可用的逻辑处理器的数量。使用这个函数，可以给每个可用的物理处理器
	在运行的时候分配一个逻辑处理器。
	PS：并不是使用多个逻辑处理器并不意味着性能更好。需配合基准测试来评估程序的运行效果
	 */
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println(runtime.NumCPU())
	/**
		WaitGroup 是一个计数信号量，可以用来记录并维护运行的goroutine
		如果WaitGroup的值大于0，Wait方法就会阻塞。
	 */
	var wg sync.WaitGroup
	/**
		表示将会有2个正在运行的goroutine
	 */
	wg.Add(2)

	fmt.Println("Start Goroutines")

	/**
		关键字defer会修改函数调用时机，在正在执行的函数返回时才真正调用defer声明的函数
	 */

	go func() {
		/**
			使用defer声明在函数退出时调用Done方法
		 */
		defer wg.Done()

		for count := 0; count < 3; count ++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	go func() {
		/**
			使用defer声明在函数退出时调用Done方法
		 */
		defer wg.Done()

		for count := 0; count < 3; count ++ {
			for char := 'A'; char < 'A'+26; char ++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	fmt.Println("Waiting to finish")
	/**
		如果WaitGroup的值大于0，Wait方法就会阻塞。
	 */
	wg.Wait()

	fmt.Println("程序结束...")

}
