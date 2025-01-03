package 基础语法

import (
	"fmt"
	"math"
)

/*
*
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
简写形式：
i := 42
f := float64(i)
u := uint(f)
*/
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
