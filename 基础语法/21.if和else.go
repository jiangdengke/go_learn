package 基础语法

import (
	"fmt"
	"math"
)

func tpow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
func main() {
	fmt.Println(
		tpow(3, 2, 10),
		tpow(3, 3, 10),
	)
}
