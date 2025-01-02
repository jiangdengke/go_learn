package main

import (
	"fmt"
	"math"
)

// math.Pow(x,n)计算x的n次方
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		fmt.Println("v", v)
		return v
	}
	return lim
}
func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
