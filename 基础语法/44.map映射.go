package 基础语法

import "fmt"

type Vertex4 struct {
	Lat, Long float64 // 定义两个字段，都是float64类型
}

var m map[string]Vertex4 //表示键是string，值是Vertex4类型

func main() {
	m = make(map[string]Vertex4) // 初始化m
	m["Bell Labs"] = Vertex4{
		40.68322, -74.233,
	} //赋值，键为"Bell Labs" ,值为Vertex4{ 40.68322, -74.233,}
	fmt.Println(m["Bell Labs"])
}
