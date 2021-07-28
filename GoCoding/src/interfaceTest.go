package main

import "fmt"

type animal interface{
	description() string
}

type cat struct {
	Type string
	Sound string
}

type snake struct {
	Type string
	Poisonous bool
}

func (s snake)description() string {
	return fmt.Sprintf("sound:%v",s.Poisonous)
}

func (c cat)description() string {
	return fmt.Sprintf("sound:%v",c.Sound)
}

func main() {
	var a animal
	a = snake{Poisonous:true}
	fmt.Println(a.description())
	a = cat{Sound:"meow!!!"}
	fmt.Println(a.description())
}

//在 main 函数中， 我们创建了一个 animal 接口类型的变量 a。我们为 animal 接口指定了 snake 和 cat 两个实例对象，并使用 Println 方法打印 a.description 。