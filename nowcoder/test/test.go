package main

import (
	"fmt"
	"sort"

	"time"
)

func main() {
	start := time.Now()
	var n, st int
	m := make(map[int][]string)
	fmt.Scan(&n)
	fmt.Scan(&st)
	score := make([]int, n)
	var name string
	for i := 0; i < n; i++ {
		fmt.Scan(&name, &score[i])
		if ns, ok := m[score[i]]; ok {
			ns = append(ns, name)
			m[score[i]] = ns
		} else {
			ns := make([]string, 0, 3)
			ns = append(ns, name)
			m[score[i]] = ns
		}
	}

	sort.Ints(score)

	exs := make(map[int]bool)

	fmt.Println()

	if st == 0 {
		for _, s := range score {
			if _, ok := exs[s]; ok {
				continue
			}
			exs[s] = true
			defer func(s int) {
				for _, name := range m[s] {
					fmt.Println(name, s)
				}
			}(s)
		}
	} else {
		for _, s := range score {
			if _, ok := exs[s]; ok {
				continue
			}
			exs[s] = true
			for _, name := range m[s] {
				fmt.Println(name, s)
			}
		}
	}
	fmt.Println(time.Now().Sub(start))
}

/*
3
1
fang 90
yang 50
ning 90
*/
