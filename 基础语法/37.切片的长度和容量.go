package 基础语法

import "fmt"

// 这里注意长度和容量的概念
// 长度就是该切片的元素的个数
// 容量是从起始位置到底层数组的末尾
func main() {
	s := []int{2, 3, 4, 6, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
