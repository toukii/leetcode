package tencent

import (
	"fmt"
	"testing"
)

// 把数组分为两组，使两组和相等
func SpiltEqual(arr []int) ([]int, []int, bool) {
	target := 0
	for _, it := range arr {
		target += it
	}
	if target%2 == 1 {
		return nil, nil, false
	}
	visited := make([]bool, len(arr))
	spiltArr := make([]int, 0, len(arr)-1)
	ret := dfs(0, 0, target>>1, arr, visited, spiltArr)
	return nil, nil, ret
}

func dfs(i, sum, target int, arr []int, visited []bool, spiltArr []int) bool {
	if sum+arr[i] > target {
		return false
	} else if sum+arr[i] == target {
		spiltArr = append(spiltArr, arr[i])
		fmt.Println(spiltArr)
		return true
	}
	visited[i] = true
	spiltArr = append(spiltArr, arr[i])
	for ii, it := range visited {
		if !it && dfs(ii, sum+arr[i], target, arr, visited, spiltArr) {
			return true
		}
	}
	spiltArr = spiltArr[:len(spiltArr)-1]
	visited[i] = false
	return false
}

func TestSE(t *testing.T) {
	arr := []int{1, 5, 4, 6, 3, 5}
	_, _, ret := SpiltEqual(arr)
	fmt.Println(ret)
	// +1 +5 +6
	// +4 +6 +5
}
