package binarytree

import (
	"fmt"
	"math"
)

type Tree struct {
	Root *TreeNode
}

func (t *Tree) Insert(key int) {
	if t.Root == nil {
		t.Root = &TreeNode{Data: key}
	} else {
		t.Root.Insert(key)
	}
}

func (t *Tree) Search(key int) bool {
	if t.Root == nil {
		return false
	}
	return t.Root.Search(key)
}

func (t *Tree) Delete(key int) {
	if t.Root == nil {
		return
	}
	t.Root.Delete(key)
}

func (t *Tree) InOrderTransversal() {
	if t.Root != nil {
		t.Root.InOrderTransversal()
	}
	fmt.Println()
}

func (t *Tree) PreOrderTransversal() {
	if t.Root != nil {
		t.Root.PreOrderTransversal()
	}
	fmt.Println()
}

func (t *Tree) PostOrderTransversal() {
	if t.Root != nil {
		t.Root.PostOrderTransversal()
	}
	fmt.Println()
}

func (t *Tree) binaryTreePaths() []string {
	if t.Root == nil {
		return []string{}
	}
	return t.Root.binaryTreePaths()
}

func (t *Tree) Min() int {
	if t.Root == nil {
		return math.MinInt
	}
	return t.Root.Min()
}

func (t *Tree) Max() int {
	if t.Root == nil {
		return math.MaxInt
	}
	return t.Root.Max()
}

func BST() {
	t := &Tree{}
	t.Insert(35)
	t.Insert(165)
	t.Insert(47)
	t.Insert(243)
	t.Insert(65)
	t.Insert(146)
	t.Insert(10)
	t.Insert(6)
	t.Insert(40)
	t.Insert(60)
	t.Insert(15)

	fmt.Println(t.Search(10))
	fmt.Println("\nInOrder Transversal: ")
	t.InOrderTransversal()
	fmt.Println("\nPreOrder Transversal: ")
	t.PreOrderTransversal()
	fmt.Println("\nPostOrder Transversal: ")
	t.PostOrderTransversal()
	fmt.Printf("\nMin: %d", t.Min())
	fmt.Printf("\nMax: %d\n", t.Max())
	t.Delete(165)
	t.InOrderTransversal()
	t.Delete(47)
	t.InOrderTransversal()
	t.Delete(15)
	t.InOrderTransversal()
	fmt.Println()

	myTree := &Tree{}
	for _, v := range []int{4, 2, 3, 0, 5} {
		myTree.Insert(v)
	}
	res := myTree.binaryTreePaths()
	fmt.Println(res)
}
