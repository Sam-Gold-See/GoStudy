package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("current i is %d\n", i)
	}
	j := 0
	for {
		if j == 5 {
			break
		}
		fmt.Printf("current j is %d\n", j)
		j++
	}

	var strAry = []string{"aa", "bb", "cc", "dd", "ee"}
	var sliceAry = make([]string, 0)
	sliceAry = strAry[1:3]
	for i, str := range sliceAry {
		fmt.Printf("current i is %d, current str is %s\n", i, str)
	}

	var dic = map[string]int{
		"apple":      1,
		"watermelon": 2,
	}
	for k, v := range dic {
		fmt.Printf("current k is %s, current v is %d\n", k, v)
	}
}
