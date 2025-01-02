package main

import "fmt"

// 提供了初始值类型可以省略，变量会从初始值中推断中该类型
var i, j int = 2, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
