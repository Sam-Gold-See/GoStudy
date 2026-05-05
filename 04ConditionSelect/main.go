package main

import "fmt"

func main() {
	localStr := "case"
	if localStr == "case" {
		fmt.Println("into true logic")
	} else {
		fmt.Println("into false logic")
	}

	var dic = map[string]int{
		"apple":      1,
		"watermelon": 2,
	}
	if num, ok := dic["orange"]; ok {
		fmt.Printf("orange num %d\n", num)
	}
	if num, ok := dic["watermelon"]; ok {
		fmt.Printf("watermelon num %d\n", num)
	}

	switch localStr {
	case "case1":
		fmt.Println("case1")
	case "case2":
		fmt.Println("case2")
	case "case3":
		fmt.Println("case3")
	default:
		fmt.Println("default")
	}
}
