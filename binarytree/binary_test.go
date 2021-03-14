package binarytree

import (
	"testing"

	. "github.com/minaorangina/data-structures-go/utils"
)

func TestBinaryTreeInsert(t *testing.T) {
	vals := []int{8, 3, 10, 1, 6, 4, 7, 14, 13}

	tree := NewBinaryTree()

	for _, v := range vals {
		tree.Insert(v)
	}

	for _, v := range vals {
		AssertTrue(t, tree.ContainsDFS(v))
	}
}

func TestBinaryTreeInOrderTraversal(t *testing.T) {
	tree := testTree()
	want := "1 3 4 6 7 8 10 13 14"

	AssertEqual(t, tree.InOrderTraversal(), want)
}

func TestBinaryTreeDFS(t *testing.T) {
	tree := testTree()

	AssertTrue(t, tree.ContainsDFS(10))
	AssertEqual(t, tree.ContainsDFS(99), false)
}
func TestBinaryTreeBFS(t *testing.T) {
	tree := testTree()

	AssertTrue(t, tree.ContainsBFS(10))
	AssertEqual(t, tree.ContainsBFS(99), false)
}

func testTree() *binaryTree {
	vals := []int{8, 3, 10, 1, 6, 4, 7, 14, 13}
	tree := NewBinaryTree()

	for _, v := range vals {
		tree.Insert(v)
	}

	return tree
}
