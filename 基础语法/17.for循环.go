package 基础语法

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 1000; i++ {
		sum += i
	}
	fmt.Println(sum)
}
