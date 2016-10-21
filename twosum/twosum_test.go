package twosum

import (
	"testing"
)

var (
	input  []int         // 输入
	filter map[int]Empty // 去除重复结果
)

func init() {
	input = []int{2, 7, 11, 15}
	filter = make(map[int]Empty)
}

type Empty struct{}

func twosum(in []int, target int) []int {
	ret := make([]int, 0, 4)
	mp := make(map[int]int)
	for i, it := range in {
		mp[it] = i
	}
	for i, it := range in {
		k := target - it
		if j, ok := mp[k]; ok {
			ret = findOne(ret, i, j)
			continue
		}
		if j, ok := mp[-k]; ok {
			ret = findOne(ret, i, j)
			continue
		}
		k = target + it
		if j, ok := mp[k]; ok {
			ret = findOne(ret, i, j)
			continue
		}
		if j, ok := mp[-k]; ok {
			ret = findOne(ret, i, j)
		}
	}
	return ret
}

func findOne(ret []int, i, j int) []int {
	if _, exist := filter[i]; exist {
		return ret
	}
	ret = append(ret, []int{i, j}...)
	filter[i] = Empty{}
	filter[j] = Empty{}
	return ret
}

func TestTwoSum(t *testing.T) {
	return
	ret := twosum(input, 9)
	t.Log(ret)
}
