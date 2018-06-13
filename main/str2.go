package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	name := "Bertking"
	fmt.Println(name)

	fmt.Println(strings.HasPrefix(name, "Bert"))
	fmt.Println(strings.HasSuffix(name, "king"))
	fmt.Println(strings.Contains(name, "k"))
	fmt.Println(strings.Index(name, "king"))
	address := "address"
	fmt.Println(strings.Index(address, "Y"))
	fmt.Println(strings.Index(address, "d"))
	fmt.Println(strings.LastIndex(address, "d"))
	fmt.Println(strings.Repeat(address, 3))

	fmt.Println(strings.Replace(address, "d", "b", 1))

	fmt.Println(address)
	/**
		对于Replace用于将字符串中的前n个字符串old替换为字符串new，并返回一个新的字符串(字符串是不能改变的)
	如果n = -1则替换所有的字符串old为字符串new
	 */
	fmt.Println(strings.Replace(address, "d", "b", -1))

	fmt.Println(strings.Count(address, "d"))

	fmt.Println(strings.ToLower(address))

	fmt.Println(strings.ToUpper(address))

	spaces := "  abc  "
	fmt.Println(len(spaces))
	fmt.Println(len(strings.TrimSpace(spaces)))
	/**
		修剪字符串
	 */
	learn := "poor english poor"
	fmt.Println(learn)
	fmt.Println(strings.TrimLeft(learn, "poor"))
	fmt.Println(strings.TrimRight(learn, "poor"))
	fmt.Println(strings.Trim(learn, "poor"))
	/**
		strings.Fields(s) 将会利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块
	 */
	sentence := "we should good good study day day up"
	fields := strings.Fields(sentence)
	fmt.Println(fields)
	/**
		strings.Split(s, sep) 用于自定义分割符号来对指定字符串进行分割
	 */
	str := "hello?world?You?learn?go"
	splits := strings.Split(str, "?")
	fmt.Println(splits)
	/**
		Join 用于将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串
	 */
	fmt.Println(strings.Join(splits, ":"))



	/**
		任何类型 T 转换为字符串总是成功的。
	 */


	/**
		返回数字 i 所表示的字符串类型的十进制数。
	 */
	fmt.Println(strconv.Itoa(10))
	fmt.Println(strconv.Itoa(-1))

}
