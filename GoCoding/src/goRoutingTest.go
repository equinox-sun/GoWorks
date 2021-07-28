package main

import (
	"fmt"
	"time"
)

func main() {
	go c()
	fmt.Println("i am main")
	time.Sleep(time.Second * 2)
}

func c() {
	time.Sleep(time.Second * 1)
	fmt.Println("I am concurrent")
}
