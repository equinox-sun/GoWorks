package main

import "fmt"

func main() {
	f()
	fmt.Println("return normally from f.")
}

func f() {
	defer func() {
		if r := recover(); //这个不懂
		r != nil {
			fmt.Println("recovered in f", r)
		}
	}()
	fmt.Println("calling g.")
	g(0)
	fmt.Println("return normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

//Defer 总是在函数结束时执行。
