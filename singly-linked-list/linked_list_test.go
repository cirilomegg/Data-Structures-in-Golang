package singly_linked_list

import "testing"

func TestLinkedListAppend(t *testing.T) {
	list := NewLinkedList()

	for i := 1; i <= 10; i++ {
		list.Append(i)
	}

	hasElement := list.Contains(4)

	if !hasElement {
		t.Errorf("Error checking if the element exists. Got %v, expected %v", hasElement, true)
	}

	if list.Head.Data != 1 {
		t.Errorf("Invalid head node. Got %v, expected %v", list.Head.Data, 1)
	}

	if list.Tail.Data != 10 {
		t.Errorf("Invalid tail node. Got %v, expected %v", list.Tail.Data, 10)
	}
}

func TestLinkedListRemove(t *testing.T) {
	list := NewLinkedList()

	for i := 1; i <= 5; i++ {
		list.Append(i)
	}

	list.Remove(5)
	hasElement := list.Contains(5)

	if hasElement {
		t.Errorf("Error checking if the node exists. Got %v, expected %v", hasElement, false)
	}

	if list.Tail.Data != 4 {
		t.Errorf("Invalid tail node. Got %v, expected %v", list.Tail.Data, 4)
	}

	list.Remove(0)
	hasElement = list.Contains(0)

	if hasElement {
		t.Errorf("Error checking if the node exists. Got %v, expected %v", hasElement, false)
	}

	if list.Head.Data != 1 {
		t.Errorf("Invalid head node. Got %v, expected %v", list.Head.Data, 1)
	}

	list.Remove(1)
	hasElement = list.Contains(1)

	if hasElement {
		t.Errorf("Error checking if the node exists. Got %v, expected %v", hasElement, false)
	}

	if list.Head.Data != 2 {
		t.Errorf("Invalid head node. Got %v, expected %v", list.Head.Data, 2)
	}

	if list.Tail.Data != 4 {
		t.Errorf("Invalid tail node. Got %v, expected %v", list.Tail.Data, 4)
	}

	list.Remove(2)

	if list.Head.Data != 3 {
		t.Errorf("Invalid head node. Got %v, expected %v", list.Head.Data, 3)
	}

	if list.Tail.Data != 4 {
		t.Errorf("Invalid tail node. Got %v, expected %v", list.Tail.Data, 4)
	}

	list.Remove(3)

	if list.Head.Data != 4 {
		t.Errorf("Invalid head node. Got %v, expected %v", list.Head.Data, 4)
	}

	list.Remove(4)

	if list.Head != nil {
		t.Errorf("An empty list should have a nil head node. Got %v, expected %v", list.Head.Data, "nil")
	}

	if list.Tail != nil {
		t.Errorf("An empty list should have a nil tail node. Got %v, expected %v", list.Tail.Data, "nil")
	}
}

func TestLinkedListReverse(t *testing.T) {
	list := NewLinkedList()

	for i := 1; i <= 10; i++ {
		list.Append(i)
	}

	list.Reverse()

	if list.Head.Data != 10 {
		t.Errorf("Invalid head node. Got %v, expected %v", list.Head.Data, 10)
	}

	if list.Tail.Data != 1 {
		t.Errorf("Invalid tail node. Got %v, expected %v", list.Head.Data, 1)
	}
}

func TestLinkedListSearch(t *testing.T) {
	list := NewLinkedList()

	for i := 1; i <= 10; i++ {
		list.Append(i)
	}

	nodeToFind := 6

	node := list.Search(nodeToFind)

	if node == nil {
		t.Errorf("Node %v not found.", nodeToFind)
	} else if node.Data != nodeToFind {
		t.Errorf("Invalid node retrieved. Got %v, expected %v", node.Data, nodeToFind)
	}

	nodeToFind = 15

	node = list.Search(nodeToFind)

	if node != nil {
		t.Errorf("Invalid node retrieved. Got %v, expected %v", node.Data, "nil")
	}
}
