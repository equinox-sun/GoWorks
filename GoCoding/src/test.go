package main

import (
	"encoding/json"
	"fmt"
	"math/cmplx"
)

func main() {
	fmt.Println("hello world")
	//变量声明
	var a bool = true
	var b int = 1
	var c string = "hello" //字符串用双引号，单引号是rune类型
	var d float32 = 1.2222
	var x complex128 = cmplx.Sqrt(-5 + 12i) //需要引入math/cmplx包，否则编译不通过
	// var x complex128 = complex(1,2) //另一种声明方法

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(x)
	fmt.Println(cmplx.Sqrt(-1)) // "(0+1i)"

	//切片
	number2 := []int{1, 2, 3, 4}
	fmt.Println(number2)
	slice1 := number2[2:]
	fmt.Println(slice1)
	slice2 := number2[:3]
	fmt.Println(slice2)
	slice3 := number2[1:4]
	fmt.Println(slice3)

	//map
	m := make(map[string]int) //key为 string 类型，value 为 int 类型
	fmt.Println(m)

	//类型转化:并不是所有类型都可以转为另一种类型。需要确保数据类型是可以转化的。
	a1 := 1.1
	b1 := int(a1)
	fmt.Println(b1)

	//指针是保存值的地址的地方。一个指针用 * 定义，根据数据类型定义指针
	var ap *int
	a2 := 12
	ap = &a2 //& 运算符可用于获取变量的地址
	fmt.Println(*ap)

	//json 内置包
	mapA := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapA) //编码
	fmt.Println(string(mapB))
	res := response{}
	json.Unmarshal(mapB, &res) //解码
	fmt.Println(res.Apple)
}

type response struct {
	Apple   int `json:"apple"`
	Lettuce int `json:"lettuce"`
}
