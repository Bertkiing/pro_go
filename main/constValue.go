package main

import (
	"fmt"
	"math"
)

/**
	omitted :忽视，忽略
 */

/**
定义的一些常量
Go鼓励不写废代码
 */
const PI = 3.1415926
const OK = true
const NOT = false


/**
在每遇到一个新的常量块或单个常量声明时， iota 都会重置为 0（ 简单地讲，每遇到一次 const 关键字，iota 就重置为 0 ）
 */
const(
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)


func main() {
	fmt.Println(calculateArea(2))
	fmt.Println(areUOk(true))

	fmt.Println(Friday)
	fmt.Println(Saturday)
}

func calculateArea(radius float64) float64 {
	return math.Pow(radius, 2) * PI
}

func areUOk(answer bool) bool {
	if answer {
		return OK
	} else {
		return NOT
	}
}
