/*
Window 环境 VS Code运行go报错：无法将“go”项识别为 cmdlet、函数、脚本文件或可运行程序的名称。
在terminal中输入 然后重启
$env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine")
*/
package main

// 声明 main 包，表明当前是一个可执行程序；如果 main 函数不在 main 包里，
// 构建工具就不会生成可执行的文件

import "fmt" // 导入内置 fmt
// import _ "github.com/goinaction/code/chapter2/sample/matchers"
// 导入的路径前面有一个下划线
// 这个技术是为了让 Go 语言对包做初始化操作，但是并不使用包里的标识符
// 为了让程序的可读性更强，Go 编译器不允许声明导入某个包却不使用
// 下划线让编译器接受这类导入，并且调用对应包内的所有代码文件里定义的 init 函数

func init() {
	fmt.Println("init")
}

func main() { // main函数，是程序执行的入口

	fmt.Println("Hello World!") // 在终端打印 Hello World!

}
