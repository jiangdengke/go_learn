package 基础语法

import "fmt"

// 切片的零值是nil
// nil切片的长度和容量为0且没有底层数组
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}

}
