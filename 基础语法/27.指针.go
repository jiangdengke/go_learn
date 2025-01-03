package 基础语法

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i         // 指向i
	fmt.Println(*p) // 拿到i对应的值：42
	*p = 21         // 通过指针p更改i对应的值，改为21
	fmt.Println(i)

	p = &j       // 指向j
	*p = *p / 37 // 更改j的值
	fmt.Println(j)
}
