package nowcoder

import (
	// "fmt"
	"testing"
)

// 最短包含字母序列长度
func ShortMatch(base, matcher string) int {
	owe := make(map[rune]int)
	sum_owe := 0
	for _, v := range matcher {
		owe[v] += 1
		sum_owe++
	}
	L, R := 0, 0
	length := len(base)
	shortest := length + 1
	for L < length && R < length {
		if sum_owe > 0 {
			if _, ok := owe[rune(base[R])]; ok {
				owe[rune(base[R])]--
				sum_owe--
			}
			R++
		} else {
			if _, ok := owe[rune(base[L])]; ok {
				owe[rune(base[L])]++
				sum_owe++
			}
			L++
		}
		if sum_owe <= 0 && shortest > R-L {
			shortest = R - L
		}
	}
	return shortest
}

func TestShortMatch(t *testing.T) {
	base := "aiibeecaeeeabrceeeeeeeeeebcd"
	matcher := "bca"
	ret := ShortMatch(base, matcher)
	t.Log(ret)
}
