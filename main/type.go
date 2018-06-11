package main

import "fmt"

/**
	Go 是一种静态类型语言：
	go 的类型：
	1.基本类型：int ,float , bool, string; (float32 , float64)
	2.复合类型：struct , array , slice ,map ,channel
	3.只描述类型的行为：interface

	PS：结构化的类型没有真正的值时，使用nil作为默认值
 */
func main() {
	var valInt = 1
	var valFloat = 3.14
	var valBool = false
	var valString = "bertking"

	var age uint
	fmt.Println(age)
	fmt.Println(valInt)
	fmt.Println(valFloat)
	fmt.Println(valBool)
	fmt.Println(valString)
	fmt.Println(5 == 6 )


	fmt.Println(isSame(1, 2))

	fmt.Println(typeTest("bertking"))

	fmt.Println("-----类型转换-----")
	aInt := 98
	fmt.Println(aInt)
	b := string(aInt) // b


	fmt.Println(b)
}

func init() {
	// 初始化包
}

/**
	函数的返回值的类型，需要写在 "函数名"&可选参数列表之后
 */
func isSame(param1 int, param2 int) bool {
	if param2 == param1 {
		return true
	} else {
		return false
	}
}

/**
	函数可以有多个返回值，返回值列表使用"()"扩起来，使用","分割
	PS：这种多返回值函数用于特定场景（如，判断某个函数执行是否成功）
 */
func typeTest(param string) (retV1 string, retV2 bool) {

	if param == "king" {
		retV1 = "ok"
		retV2 = true
		return retV1, retV2
	} else if param == "bertking" {
		retV1 = "yes"
		retV2 = true
	} else {
		retV1 = "holyhigh"
		retV2 = true
	}
	return retV1, retV2
}
