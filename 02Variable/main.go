package main

import "fmt"

var globalStr string
var globalInt int

var GlobalStr string
var GlobalInt int

// 判断是否包外可见根据首字母大小写情况

func main() {
	var localStr string
	var localInt int
	localStr = "first local"
	localInt = 2026
	globalInt = 1024
	globalStr = "first global"

	fmt.Printf("globalStr is %s\n", globalStr)
	fmt.Printf("globalInt is %d\n", globalInt)
	fmt.Printf("localStr is %s\n", localStr)
	fmt.Printf("localInt is %d\n", localInt)
}
