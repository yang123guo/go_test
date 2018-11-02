// 当前程序的包名称
package main

// 导入其他的包
import (
    "fmt"
    // "io"
    // "os"
    // "time"
)

// 常量的定义
const PI = 3.14

const (
    const1 = "1"
    const2 = 2
)

// 全部变量的申明和赋值
var name = "faith"

// 一般类型申明
type newType int

// 结构的申明
type gopher struct{}

// 接口的申明
type golang interface{}

// 由main函数作为程序的入口点启动
func main() {
    fmt.Println("hello faith")
}
