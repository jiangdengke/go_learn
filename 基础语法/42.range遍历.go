package 基础语法

import "fmt"

// for循环的range形式可遍历切片或映射
// 使用for循环遍历切片时，每次迭代都会返回两个值，第一个值为当前元素的下标，第二个值为该下标所对应的一份副本

var pow1 = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow1 {
		fmt.Print(i, v)
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
