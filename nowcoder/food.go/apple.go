package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sort.Ints(a)

	fmt.Println(do(a, n))
}

func do(a []int, n int) int {
	sum := 0
	jiou := a[0] % 2
	for _, it := range a {
		sum += it
		if it%2 != jiou {
			return -1
		}
	}
	avg := sum / n
	// fmt.Println(avg, sum%n)
	if sum%n != 0 {
		return -1
	}

	avgPos := 0
	for i := 0; i < n-1; i++ {
		if a[i] <= avg && a[i+1] >= avg {
			avgPos = i
			break
		}
	}
	ret := 0
	i, j := avgPos, avgPos+1
	for i >= 0 || j < n {
		for j < n && a[j] == avg {
			j++
		}
		for i >= 0 && a[i] == avg {
			i--
		}
		if i < 0 || j >= n {
			break
		}
		m := min(avg, a[i], a[j])
		// fmt.Println("#", a, i, j)
		if m == 0 {
			ret += (a[j] - avg) / 2
			a[i], a[j] = avg, avg
			i--
			j++
		} else if m > 0 {
			ch := avg - a[i]
			ret += (ch) / 2
			a[j] -= ch
			a[i] = avg
			i--
		} else {
			ch := a[j] - avg
			ret += (ch) / 2
			a[i] += ch
			a[j] = avg
			j++
		}
		// fmt.Println("@", a, i, j)
	}
	return ret
}

func min(avg, m1, m2 int) int {
	ch1, ch2 := avg-m1, m2-avg
	if ch1 == ch2 {
		return 0
	} else if ch1 > ch2 {
		return -ch2 // 按正数，右移动
	} else {
		return ch1 // 按负数,左移动
	}
	return 0
}
