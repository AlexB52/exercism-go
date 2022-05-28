package binarysearchtree

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
	if bst == nil {
		return
	}

	return append(append(bst.left.SortedData(), bst.data), bst.right.SortedData()...)
}
