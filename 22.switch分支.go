package main

import (
	"fmt"
	"runtime"
)

// GO的switch与java不同，它默认在每个case后都加了一个break
func main() {
	fmt.Print("GO 运行的系统环境：")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}
