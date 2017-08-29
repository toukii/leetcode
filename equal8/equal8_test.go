package equal8

import (
	"fmt"
	"testing"
)

func Equal8(input []int) (i, j int) {
	size := len(input)
	if size < 2 {
		return 0, 0
	}
	j = size - 1
	for input[j] >= 8 {
		j--
	}
	for i < j {
		fmt.Println(input[i], input[j])
		if input[i]+input[j] == 8 {
			return input[i], input[j]
		} else if input[i]+input[j] > 8 {
			j--
		} else {
			i++
		}
	}
	return 0, 0
}

func TestEqual8(t *testing.T) {
	input := []int{1, 3, 4, 5, 6, 7, 8, 9, 10}
	i, j := Equal8(input)
	t.Logf("i:%d, j:%d\n", i, j)

	input1 := []int{1, 3, 4, 6, 6, 8, 8, 9, 10}
	i, j = Equal8(input1)
	t.Logf("i:%d, j:%d\n", i, j)
}
