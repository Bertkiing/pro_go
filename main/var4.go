package main

import (
	"fmt"
	"runtime"
	"os"
)

/**
	这种因式分解的写法一般用于声明全局变量
 */
var (
	name    = "Bertking"
	age     = 25
	address = "LY"
)


func main() {
	testVarInit()
	fmt.Printf("Name is %s\n", name)
	fmt.Printf("Age is %d\n", age)
	fmt.Printf("Address is %s\n", address)

	var goos = runtime.GOOS
	fmt.Printf("The operating system is : %s\n ", goos)
	/**
		当你在函数体内声明局部变量时，应使用简短声明语法 :=
	 */
	path := os.Getenv("PATH")

	fmt.Printf("Path is %s \n", path)

}

/**
	当变量被声明之后，系统自动赋予它该类型的初始值:
	int - 0
	float - 0.0
	bool - false
	string - ""
	指针 - nil
	复合类型 - 对应的
 */
func testVarInit() {
	var gender string
	var assets int
	var radius float64
	var isMarried bool
	var pointer *int
	var arr [2]int  // [0,0]
	var slice []int // []

	fmt.Println(gender == "") // 声明string 的变量，默认值是空字符串 ""
	fmt.Println(assets)
	fmt.Println(radius)
	fmt.Println(isMarried)
	fmt.Println(pointer)
	fmt.Println(arr)
	fmt.Println(slice)
	// 空数组和空切片有区别吗？？？
}
