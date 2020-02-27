package main

import (
	"strconv"
	"fmt"
)

const name_test = "Hello,Go"
const x, y int = 1, 2

const (
	xx, yy = 1, 2
	isTrue = false
)

func main() {
	x, _ := strconv.Atoi("12")
	println(x)
	fmt.Println("Hello Go...")

	const mobile = 10086
	const email  = "xxx@163.com"
	println(email)
}
