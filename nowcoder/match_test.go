package nowcoder

import (
	"testing"
)

func DeepMatch(s, e string, si, ei int) bool {
	sL := len(s)
	eL := len(e)
	if ei >= eL {
		return si >= sL
	}
	// e[ei+1]!=*
	if ei+1 >= eL || '*' != e[ei+1] {
		return (s[si] == e[ei] || '.' == e[ei]) && DeepMatch(s, e, si+1, ei+1)
	}
	// e[ei+1] == *
	for si < sL && (s[si] == e[ei] || '.' == e[ei]) {
		if DeepMatch(s, e, si, ei+2) {
			return true
		}
		si++
	}
	// ignore X*: treate X* as ""
	return DeepMatch(s, e, si, ei+2)
}

func TestMatch(t *testing.T) {
	s := "aabc"
	e := "a*b.*"
	matched := DeepMatch(s, e, 0, 0)
	t.Log(matched)
}
