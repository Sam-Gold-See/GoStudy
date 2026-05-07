package main

import (
	"fmt"
	"time"
)

func a() {
	time.Sleep(3 * time.Second)
	fmt.Println("it's a")
}

func b() {
	time.Sleep(2 * time.Second)
	fmt.Println("it's b")
}
func c() {
	time.Sleep(1 * time.Second)
	fmt.Println("it's c")
}

func main() {
	// 使用 go fun 开启协程
	go a()
	go b()
	go c()
	time.Sleep(5 * time.Second)
}
