package 基础语法

import "fmt"

func main() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

// 这样的话，就是defer把数压到了一个栈中。然后9才是先出
/*
counting
done
9
8
7
6
5
4
3
2
1
0
*/
