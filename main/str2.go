package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Bertking"
	fmt.Println(name)

	fmt.Println(strings.HasPrefix(name,"Bert"))
	fmt.Println(strings.HasSuffix(name,"king"))
	fmt.Println(strings.Contains(name,"k"))
	fmt.Println(strings.Index(name,"king"))
	address := "address"
	fmt.Println(strings.Index(address,"Y"))
	fmt.Println(strings.Index(address,"d"))
	fmt.Println(strings.LastIndex(address,"d"))
	fmt.Println(strings.Repeat(address,3))

	fmt.Println(strings.Replace(address,"d","b",1))

	fmt.Println(address)
	/**
		对于Replace用于将字符串中的前n个字符串old替换为字符串new，并返回一个新的字符串(字符串是不能改变的)
	如果n = -1则替换所有的字符串old为字符串new
	 */
	fmt.Println(strings.Replace(address,"d","b",-1))

	fmt.Println(strings.Count(address,"d"))

	fmt.Println(strings.ToLower(address))

	fmt.Println(strings.ToUpper(address))


	spaces := "  abc  "
	fmt.Println(len(spaces))
	fmt.Println(len(strings.TrimSpace(spaces)))

	learn := "poor english poor"
	fmt.Println(learn)
	fmt.Println(strings.TrimLeft(learn,"poor"))
	fmt.Println(strings.TrimRight(learn,"poor"))
	fmt.Println(strings.Trim(learn,"poor"))


}
