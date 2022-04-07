// main.go

package main

import (
    "fmt"
)

// 必须是 main 包：package main
// 必须是 main ⽅法：func main()
// ⽂件名不⼀定是 main.go

// Go 中 main 函数不⽀持任何返回值
// 通过 os.Exit 来返回状态

func main() {
    fmt.Println("Hello world!")
}
