package merge_sort

/*
https://pkg.go.dev/container/list@go1.17.5#Element
 */

import (
	"github.com/panutat-p/data-structures-algorithms-beginner/linked_list"
)

// MergeSortLinkedList O(kn log n) time
func MergeSortLinkedList(li linked_list.LinkedList) linked_list.LinkedList {
	if li.Size() == 1 {  // base case
		return li
	}

	left, right := SplitLinkedList(li)
	left = MergeSortLinkedList(left)
	right = MergeSortLinkedList(right)

	return MergeLinkedList(left, right)
}

// SplitLinkedList O(k log n) time
func SplitLinkedList(li linked_list.LinkedList) (linked_list.LinkedList, linked_list.LinkedList) {
	size := li.Size()

	if size == 0 {
		return li, linked_list.LinkedList{}
	}

	// midpoint is a node at index size/2
	// traverse to the node before midpoint
	// because we want to assign nil to the end of the left linked list
	index := (size / 2) - 1
	current := li.Head
	for i := 0; i < index; i += 1 {
		current = current.NextNode
	}

	right := linked_list.LinkedList{current.NextNode}
	current.NextNode = nil  // cut the chain at the midpoint
	return li, right
}

// MergeLinkedList O(n) time
func MergeLinkedList(left linked_list.LinkedList, right linked_list.LinkedList) linked_list.LinkedList {
	n := linked_list.Node{0, nil}  // fake node
	li := linked_list.LinkedList{&n}

	current := li.Head
	leftNode := left.Head
	rightNode := right.Head
	for leftNode != nil || rightNode != nil {
		if leftNode == nil {  // end of left linked list
			current.NextNode = rightNode  // append the right linked list
			break
		} else if rightNode == nil {  // end of right linked list
			current.NextNode = leftNode  // append the left linked list
			break
		} else {  // normal case
			if leftNode.Data < rightNode.Data {
				current.NextNode = leftNode  // append the left linked list
				leftNode = leftNode.NextNode
			} else {
				current.NextNode = rightNode  // append the right linked list
				rightNode = rightNode.NextNode
			}
		}
		current = current.NextNode
	}

	li.Head = li.Head.NextNode  // remove fake node
	return li
}
