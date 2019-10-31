package main

import "fmt"

//简单来讲：go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。
//示例代码 main函数相当
func main() {
	fmt.Println("c")
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("d")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
		}
		fmt.Println("e")
	}()
	f()              //开始调用f
	fmt.Println("f") //这里开始下面代码不会再执行
}

func f() {
	fmt.Println("a")
	panic("异常信息")    //抛出异常后，不在向下执行，defer方法中去处理异常
	fmt.Println("b") //这里开始下面代码不会再执行
}

//-------output-------
//c
//a
//d
//异常信息
//e
