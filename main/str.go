package main

import "fmt"

func main() {
	var name  = "bertking"
	fmt.Println(name)
	fmt.Printf("name' length %d\n",len(name))// 获取字符串长度
	fmt.Printf("name's first is %c\n",name[0]) // 获取字符串的第一个字节
	fmt.Printf("name's last is %c\n",name[len(name)-1])

	// 不能获取字符串中某个字节的地址 cannot take the address of name[0]
	// fmt.Println(&name[0])

	sayHello(name) // 字符串拼接


	/**
		千万要注意：这里使用的是反引号
	 */
	poem := `九霄龙云惊天变,
             风云际会浅水游`

	fmt.Println(poem)

	test := `Next station \n` // 转义符原样输出
	fmt.Println(test)
}


func sayHello(name string)  {
	fmt.Println("hello,"+name)
}



