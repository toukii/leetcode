/*查找单词，是否存在*/
package huawei

import (
	"fmt"
	"testing"
)

var (
	m, n    int
	letters [][]string
	visited [][]bool
	target  string
)

func init() {
	m, n = 5, 5
	letters = [][]string{
		{"C", "P", "U", "C", "Y"},
		{"E", "K", "L", "Q", "H"},
		{"C", "R", "S", "O", "L"},
		{"F", "A", "I", "A", "O"},
		{"P", "G", "R", "B", "C"},
	}
	visited = [][]bool{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}
	target = "SOLO"
}

func TestWord(t *testing.T) {
	fmt.Println(letters)
	fmt.Println(visited)
	found := FindWord()
	fmt.Println(found)
}

func FindWord() bool {
	start := string(target[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if letters[i][j] == start {
				found := findFromHere(i, j, 0)
				if found {
					return true
				}
			}
		}
	}
	return false
}

func findFromHere(i, j, th int) bool {
	if visited[i][j] {
		return false
	}
	if th+1 >= len(target) {
		return true
	}
	visited[i][j] = true
	var up, down, left, right bool
	cur := string(target[th+1])
	if i > 0 && cur == letters[i-1][j] && !visited[i-1][j] {
		up = findFromHere(i-1, j, th+1)
	}
	if i < m-1 && cur == letters[i+1][j] && !visited[i+1][j] {
		down = findFromHere(i+1, j, th+1)
	}
	if j > 0 && cur == letters[i][j-1] && !visited[i][j-1] {
		left = findFromHere(i, j-1, th+1)
	}
	if j < n-1 && cur == letters[i][j+1] && !visited[i][j+1] {
		down = findFromHere(i, j+1, th+1)
	}
	if up || down || left || right {
		return true
	}
	visited[i][j] = false
	return false
}
