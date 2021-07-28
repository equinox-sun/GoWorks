package main
import "fmt"


//struct
type person struct{
	name string
	age int
	gender string
}

//define function
//返回值是指针类型
//how many print function:Printf/Println
func (p *person) describe() {
	fmt.Printf("%v is %v years old.",p.name,p.age)
}

func (p *person) setAge(age int) {
	p.age=age
}

func (p person) setName(name string) {
	p.name=name
}
//这两个方法的区别：第一种返回值是类型指针，更改了函数外部的age。第二种返回值是值类型，不改变外部值。说白了，值传递与地址传递的区别

func main() {
	pp := &person{name:"Bob",age:42,gender:"Male"}
	pp.describe()
	pp.setAge(45)
	fmt.Println(pp.age)
	pp.setName("hari")
	fmt.Println(pp.name)//Bob
}