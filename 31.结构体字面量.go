package main

import "fmt"

type Vertex3 struct {
	X, Y int
}

var (
	v1 = Vertex3{1, 2}  // 创建一个Vertex3类型的结构体并赋值
	v2 = Vertex3{X: 1}  // Y:0 被隐式的赋予零值
	v3 = Vertex3{}      // X:0 Y:0
	p  = &Vertex3{1, 2} // 创建一个 *Vertex3类型的结构体（指针）
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
