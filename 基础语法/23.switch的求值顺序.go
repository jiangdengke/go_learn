package 基础语法

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("周六是哪天？")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today:
		fmt.Println("今天")
	case today + 1:
		fmt.Println("明天")
	case today + 2:
		fmt.Println("后天")
	default:
		fmt.Println("很多天后")
	}
}
