package alibaba

// 如数组：[1,3,3,1,3],找出重复次数为3的数字(其它数字重复次数为2次).
// 要求时间复杂度不大于O(nlogn),空间复杂度为O(1).

import (
	"sort"
	"testing"
)

func find(isli []int) int {
	length := len(isli)
	if length < 3 {
		panic("length error.")
	}
	v := isli[0]
	for i := 1; i < length; i++ {
		v ^= isli[i]
	}
	return v
}

func find2(isli []int) int {
	// quickSort
	sort.Ints(isli)
	times := 1
	cur := isli[0]
	length := len(isli)
	for i := 1; i < length; i++ {
		if cur == isli[i] {
			times++
			if times == 3 {
				break
			}
		} else {
			times = 1
			cur = isli[i]
		}
	}
	return cur
}

func find3(isli []int) int {
	// quickSort
	sort.Ints(isli)
	length := len(isli)
	for i := 2; i < length; i += 2 {
		if isli[i-2] == isli[i] {
			return isli[i]
		}
	}
	return -1
}

func TestFind(t *testing.T) {
	// isli := []int{1, 3, 3, 1, 3}
	isli := []int{1, 3, 3, 1, 3, 5, 5, 6, 6}
	isli2 := make([]int, 5)
	copy(isli2, isli)

	ret := find(isli)
	t.Log(ret)

	ret2 := find2(isli[:])
	t.Log(ret2)

	ret3 := find3(isli2)
	t.Log(ret3)

}
