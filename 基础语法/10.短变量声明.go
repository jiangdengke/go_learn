package 基础语法

import "fmt"

// 在函数中，短赋值语句：=可在隐式确定类型的var声明中使用。因此：=结构不能在函数外使用
func main() {
	var i = 1
	var j = 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
