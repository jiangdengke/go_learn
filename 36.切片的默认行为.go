package main

import "fmt"

func main() {
	s := []int{2, 3, 4, 5, 6}
	s = s[1:4]
	fmt.Println(s)

	// 省略起始索引时，默认为0
	s = s[:2]
	fmt.Println(s)

	// s[1:]是对上一个s切片。终止索引为数组的长度
	s = s[1:]
	fmt.Println(s)
}
