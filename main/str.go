package main

import "fmt"

func main() {
	var name string = "bertking"
	fmt.Println(name)
	fmt.Printf("name' length %d\n",len(name))
	fmt.Printf("name's first is %c\n",name[0])
	fmt.Printf("name's last is %c\n",name[len(name)-1])

	/**
		千万要注意：这里使用的是反引号
	 */
	poem := `九霄龙云惊天变,
             风云际会浅水游`

	fmt.Println(poem)

	test := `Next station \n` // 转义符原样输出
	fmt.Println(test)




}
