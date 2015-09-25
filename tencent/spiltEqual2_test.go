package tencent

import (
	"fmt"
	"testing"
)

// 把数组分为两组，使两组加减相等
func SpiltEqual2(arr []int) bool {
	Open := make(map[int]bool)
	length := len(arr)
	for i := 0; i < length; i++ {
		Open[i] = true
	}
	spiltArr := make([]int, 0, length)
	Open[length-1] = false
	delete(Open)
	ret := DFS(0, 0, arr, Open, spiltArr, true) || DFS(0, 0, arr, Open, spiltArr, false)
	return ret
}

func DFS(i, sum int, arr []int, Open map[int]bool, spiltArr []int, flag bool) bool {
	cur := arr[i]
	if !flag {
		cur *= -1
	}
	if condition {

	}
	if sum+cur == target {
		spiltArr = append(spiltArr, cur)
		fmt.Println(spiltArr)
		return true
	}
	visited[i] = true
	spiltArr = append(spiltArr, cur)
	for ii, it := range visited {
		if !it && (DFS(ii, sum+cur, target, arr, visited, spiltArr, true) || DFS(ii, sum+cur, target, arr, visited, spiltArr, false)) {
			return true
		}
	}
	spiltArr = spiltArr[:len(spiltArr)-1]
	visited[i] = false
	return false
}

func TestSE2(t *testing.T) {
	arr := []int{1, 5, -4, 6, 3, 5}
	_, _, ret := SpiltEqual2(arr)
	fmt.Println(ret)
	// +1 +5 +6
	// +4 +6 +5
}
