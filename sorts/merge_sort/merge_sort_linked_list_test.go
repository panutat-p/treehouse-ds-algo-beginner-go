package merge_sort

import (
	"fmt"
	"github.com/panutat-p/data-structures-algorithms-beginner/linked_list"
	"testing"
)

func TestSplitLinkedListOdd(t *testing.T) {
	li := linked_list.LinkedList{}
	n3 := linked_list.Node{3, nil}
	n2 := linked_list.Node{2, &n3}
	n1 := linked_list.Node{1, &n2}
	li.Head = &n1

	left, right := SplitLinkedList(li)
	left.Traverse()
	right.Traverse()

	if left.Size() != 1 {
		t.Errorf("got %v, but expect %v", left.Size(), 1)
	}

	if right.Size() != 2 {
		t.Errorf("got %v, but expect %v", right.Size(), 2)
	}
}

func TestSplitLinkedListEven(t *testing.T) {
	li := linked_list.LinkedList{}
	n4 := linked_list.Node{4, nil}
	n3 := linked_list.Node{3, &n4}
	n2 := linked_list.Node{2, &n3}
	n1 := linked_list.Node{1, &n2}
	li.Head = &n1
	li.Traverse()

	left, right := SplitLinkedList(li)
	fmt.Print("left: ")
	left.Traverse()
	fmt.Print("right: ")
	right.Traverse()

	if left.Size() != 2 {
		t.Errorf("got %v, but expect %v", left.Size(), 2)
	}

	if right.Size() != 2 {
		t.Errorf("got %v, but expect %v", right.Size(), 2)
	}
}

func TestMergeLinkedList(t *testing.T) {
	left := linked_list.LinkedList{}
	n3 := linked_list.Node{7, nil}
	n2 := linked_list.Node{5, &n3}
	n1 := linked_list.Node{1, &n2}
	left.Head = &n1

	right := linked_list.LinkedList{}
	n33 := linked_list.Node{8, nil}
	n22 := linked_list.Node{4, &n33}
	n11 := linked_list.Node{2, &n22}
	right.Head = &n11

	li := MergeLinkedList(left, right)
	li.Traverse()

	if li.Size() != 6 {
		t.Errorf("got %v, but expect %v", li.Size(), 6)
	}
}

func TestMergeSortLinkedList(t *testing.T) {
	li := linked_list.LinkedList{}
	n7 := linked_list.Node{1, nil}
	n6 := linked_list.Node{50, &n7}
	n5 := linked_list.Node{99, &n6}
	n4 := linked_list.Node{9, &n5}
	n3 := linked_list.Node{34, &n4}
	n2 := linked_list.Node{2, &n3}
	n1 := linked_list.Node{100, &n2}
	li.Head = &n1
	li.Traverse()

	sortedList := MergeSortLinkedList(li)
	sortedList.Traverse()

	if sortedList.Size() != 7 {
		t.Errorf("got %v, but expect %v", sortedList.Size(), 7)
	}
}
