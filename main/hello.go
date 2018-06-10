/**
	可以通过该程序，证实一点：Go程序先执行init(),后执行main()函数
 */
package main

import "fmt"

func main() {
	fmt.Println("Hello,Go")
}

func init(){
	fmt.Println("hello , init()")
}
