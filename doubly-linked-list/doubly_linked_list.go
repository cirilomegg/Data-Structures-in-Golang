package doubly_linked_list

import (
	"log"
	"strconv"
	"strings"
)

type Node struct {
	Data     int
	Next     *Node
	Previous *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func NewLinkedList() LinkedList {
	return LinkedList{
		Head: nil,
	}
}

func (list *LinkedList) Append(data int) {
	element := &Node{
		Data: data,
	}

	if list.Head == nil {
		list.Head = element
		list.Tail = element
		return
	}

	current := list.Head

	for current != nil {
		if current.Next == nil {
			element.Previous = current
			current.Next = element
			list.Tail = element
			return
		}

		current = current.Next
	}
}

func (list *LinkedList) Contains(data int) bool {
	node := list.Search(data)
	return node != nil
}

func (list *LinkedList) Search(data int) *Node {
	current := list.Head

	for current != nil {
		if current.Data == data {
			return current
		}

		current = current.Next
	}

	return nil
}

func (list *LinkedList) Remove(data int) {
	if list == nil {
		return
	}

	current := list.Head
	var previous *Node

	for current != nil {
		if current.Data == data {
			if current == list.Head {
				list.Head = current.Next

				if list.Head != nil {
					list.Head.Previous = nil
				}
			} else {
				previous.Next = current.Next
				current.Previous = previous
			}

			if list.Tail == current {
				list.Tail = current.Previous
			}

			current = nil
			return
		}

		previous = current
		current = current.Next
	}
}

func (list *LinkedList) Reverse() {
	list.Tail = list.Head
	current := list.Head
	var previous *Node
	var next *Node

	for current != nil {
		next = current.Next
		current.Previous = current.Next
		current.Next = previous
		previous = current
		current = next
	}

	list.Head = previous
}

func (list *LinkedList) Print() {
	if list == nil || list.Head == nil {
		return
	}

	builderForward := strings.Builder{}
	builderBackward := strings.Builder{}

	current := list.Head

	for current != nil {
		builderForward.WriteString(strconv.Itoa(current.Data))

		if current.Next != nil {
			builderForward.WriteString(" -> ")
		}

		current = current.Next
	}

	current = list.Tail

	for current != nil {
		builderBackward.WriteString(strconv.Itoa(current.Data))

		if current.Previous != nil {
			builderBackward.WriteString(" -> ")
		}

		current = current.Previous
	}

	log.Println(builderForward.String())
	log.Println(builderBackward.String())
}
