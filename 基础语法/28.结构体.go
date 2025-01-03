package 基础语法

import "fmt"

type Vertex struct {
	X int
	Y int
}

// 一个 结构体（struct）就是一组 字段（field）。
func main() {
	fmt.Println(Vertex{1, 2})

}
