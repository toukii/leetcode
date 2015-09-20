package nowcoder

import (
	"fmt"
	"testing"
)

// 所有都是正数，求固定和最长字序列的长度
func LongSum(arr []int, target int) int {
	length := len(arr)
	sum := 0
	Long := 0
	L, R := 0, 0
	for L < length && R < length {
		if sum < target {
			sum += arr[R]
			R++
		} else if sum > target {
			sum -= arr[L]
			L++
		} else {
			if R-L > Long {
				Long = R - L
			}
			sum -= arr[L]
			L++
		}
	}
	if sum == target && R-L > Long {
		return R - L
	}
	return Long
}

func TestLongSum(t *testing.T) {
	arr := []int{1, 2, 1, 1, 1, 2}
	target := 3
	ret := LongSum(arr, target)
	if ret != 3 {
		t.Error(ret)
	} else {
		t.Log(ret)
	}
}
