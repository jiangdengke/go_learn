package 基础语法

import "fmt"

// go可以返回多个返回值
func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	fmt.Println(swap("hello", "world"))
}
