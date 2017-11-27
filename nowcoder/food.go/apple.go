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

/*
题目描述
n 只奶牛坐在一排，每个奶牛拥有 ai 个苹果，现在你要在它们之间转移苹果，使得最后所有奶牛拥有的苹果数都相同，每一次，你只能从一只奶牛身上拿走恰好两个苹果到另一个奶牛上，问最少需要移动多少次可以平分苹果，如果方案不存在输出 -1。
输入描述:
每个输入包含一个测试用例。每个测试用例的第一行包含一个整数 n（1 <= n <= 100），接下来的一行包含 n 个整数 ai（1 <= ai <= 100）。
输出描述:
输出一行表示最少需要移动多少次可以平分苹果，如果方案不存在则输出 -1。
示例1
输入

4
7 15 9 5
输出
3
*/
