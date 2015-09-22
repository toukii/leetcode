package tencent

import (
	"testing"
)

func SpiltEqual(arr []int) ([]int, []int, bool) {
	return nil, nil, false
}

func dfs() {

}

func TestSE(t *testing.T) {
	arr := []int{1, 5, 2, 6, 3, 5}
	SpiltEqual(arr)
	// 1+2+3-6+5-5 = 0
	// 1,2,3,5  == 5,6
}
