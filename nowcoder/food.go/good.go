package main

import (
	"fmt"
)

func main() {
	var input string
	empty := struct{}{}
	m := make(map[string]struct{})
	var n int
	for {
		n, _ = fmt.Scanf("%s", &input)
		// fmt.Println(input, n, err)
		m[input] = empty
		// fmt.Println(len(m), m)
		if n <= 0 {
			break
		}
	}
	fmt.Println(len(m))
}

// 运行时间：7ms
// 占用内存：752k
