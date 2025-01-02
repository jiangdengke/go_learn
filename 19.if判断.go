package main

import (
	"fmt"
	"math"
)

// math.sqrt是计算平方跟
// fmt.Sprint是将结果转化为字符串输出
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func main() {
	fmt.Println(sqrt(2), sqrt(-4))

}
