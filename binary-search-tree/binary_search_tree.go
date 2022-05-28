package binarysearchtree

// import "fmt"

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

func (bst *BinarySearchTree) Insert(i int) {
	if i > bst.data {
		if bst.right != nil {
			bst.right.Insert(i)
		} else {
			bst.right = &BinarySearchTree{left: bst.right, data: i}
		}
	} else {
		if bst.left != nil {
			bst.left.Insert(i)
		} else {
			bst.left = &BinarySearchTree{right: bst.left, data: i}
		}
	}
}

func (bst *BinarySearchTree) SortedData() (result []int) {
	if bst.left != nil {
		result = append(result, bst.left.SortedData()...)
	}

	result = append(result, bst.data)

	if bst.right != nil {
		result = append(result, bst.right.SortedData()...)
	}

	return result
}
