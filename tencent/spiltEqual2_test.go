package tencent

import (
	"fmt"
	"testing"
)

var (
	times = 0
)

// 数组元素使用加减法，和为0
func SpiltEqual2(arr []int) bool {
	Open := make(map[int]bool)
	length := len(arr)
	for i := 0; i < length; i++ {
		Open[i] = true
	}
	spiltArr := make([]int, 0, length)
	ret := DFS(0, 0, arr, Open, spiltArr, true) || DFS(0, 0, arr, Open, spiltArr, false)
	return ret
}

func DFS(i, sum int, arr []int, Open map[int]bool, spiltArr []int, flag bool) bool {
	times++
	fmt.Print(times, ".")
	cur := arr[i]
	if !flag {
		cur *= -1
	}
	delete(Open, i)
	sum += cur
	spiltArr = append(spiltArr, cur)

	if len(Open) <= 0 && 0 == sum {
		fmt.Println(spiltArr)
		return true
	}
	for ii, _ := range Open {
		if DFS(ii, sum, arr, Open, spiltArr, true) || DFS(ii, sum, arr, Open, spiltArr, false) {
			return true
		}
	}

	spiltArr = spiltArr[:len(spiltArr)-1]
	Open[i] = true
	return false
}

func TestSE2(t *testing.T) {
	arr := []int{1, 5, -4, 6, 3, 5}
	ret := SpiltEqual2(arr)
	fmt.Println(ret)
	// +1 +5 +6
	// -4 -3 -5
}
