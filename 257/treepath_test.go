package _257_test

import (
	"fmt"
	"testing"
	"time"
)

type TreeNode struct {
	v           int
	left, right *TreeNode
}

func NewTreeNode(v int, l, r *TreeNode) *TreeNode {
	return &TreeNode{
		v:     v,
		left:  l,
		right: r,
	}
}

var result chan string

func init() {
	result = make(chan string, 10)
}

func binaryTreePaths(root *TreeNode, v string) {
	if nil == root {
		return
	}
	v += fmt.Sprintf("-%d", root.v)
	if root.left == nil && root.right == nil {
		result <- v
		return
	}
	binaryTreePaths(root.left, v)
	binaryTreePaths(root.right, v)
}

func _main() {
	c1 := NewTreeNode(5, nil, nil)
	b1 := NewTreeNode(2, nil, c1)
	b2 := NewTreeNode(3, nil, nil)
	a1 := NewTreeNode(1, b1, b2)
	go func() {
		for {
			select {
			case v := <-result:
				fmt.Println(v)
			}
		}
	}()
	binaryTreePaths(a1, "")
	time.Sleep(1e9)
}

func TestPath(t *testing.T) {
	_main()
}
