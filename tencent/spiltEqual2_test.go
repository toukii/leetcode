package tencent

import (
	"fmt"
	"testing"
)

// 把数组分为两组，使两组加减相等
func SpiltEqual2(arr []int) ([]int, []int, bool) {
	target := 0
	for _, it := range arr {
		target += it
	}
	if target%2 == 1 {
		return nil, nil, false
	}
	visited := make([]bool, len(arr))
	spiltArr := make([]int, 0, len(arr)-1)
	ret := DFS(0, 0, target>>1, arr, visited, spiltArr, true)
	return nil, nil, ret
}

func DFS(i, sum, target int, arr []int, visited []bool, spiltArr []int, flag bool) bool {
	cur := arr[i]
	if !flag {
		cur *= -1
	}
	if sum+arr[i] > target {
		return false
	} else if sum+cur == target {
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
