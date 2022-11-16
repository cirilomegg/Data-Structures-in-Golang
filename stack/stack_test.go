package stack

import "testing"

func TestNewStack(t *testing.T) {
	stack := NewStack()

	if !stack.IsEmpty() {
		t.Errorf("A newly instantiated stack should have no elements. Got %v, expected %v", stack.Length(), 0)
	}

	if stack.Length() != 0 {
		t.Errorf("A newly instantiated stack should have no elements. Got %v, expected %v", stack.Length(), 0)
	}
}

func TestStackPush(t *testing.T) {
	stack := NewStack()
	maxLength := 10

	for i := 1; i <= maxLength; i++ {
		stack.Push(i)
	}

	if stack.Top.Data != maxLength {
		t.Errorf("Error pushing first element to stack. Got %v, expected 10 on the 'tail' element", stack.Top.Data)
	}

	current := stack.Top
	i := maxLength
	for current != nil {
		if current.Data != i {
			t.Errorf("Error pushing element to stack. Got %v, expected %v", current.Data, i)
		}

		current = current.Next
		i--
	}
}

func TestStackLength(t *testing.T) {
	stack := NewStack()
	stackLength := 10

	for i := 1; i <= stackLength; i++ {
		stack.Push(i)
	}

	if stack.Length() != stackLength {
		t.Errorf("Wrong stack length. Got %v, expected %v", stack.Length(), stackLength)
	}
}

func TestStackIsEmpty(t *testing.T) {
	stack := NewStack()

	if !stack.IsEmpty() {
		t.Errorf("Stack should be empty. Got %v, expected %v", stack.Length(), 0)
	}

	stack.Push(1)

	if stack.IsEmpty() {
		t.Errorf("Stack should not be empty. Got %v, expected %v", stack.Length(), 0)
	}
}

func TestStackPeek(t *testing.T) {
	stack := NewStack()

	_, err := stack.Peek()

	if err == nil {
		t.Errorf("Peeking an empty stack should return an error.")
	}

	stackLength := 10

	for i := 1; i <= stackLength; i++ {
		stack.Push(i)

		item, err := stack.Peek()

		if err != nil {
			t.Errorf("Peeking a populated stack should not return an error")
		}

		if item != i {
			t.Errorf("Error peeking element from stack. Got %v, expected %v", item, i)
		}
	}

	currentLength := stack.Length()

	if currentLength != stackLength {
		t.Errorf("Peeking the top element from stack should not decrease the length. Got %v, expected %v", currentLength, stackLength)
	}
}

func TestStackPop(t *testing.T) {
	stack := NewStack()

	_, err := stack.Pop()

	if err == nil {
		t.Errorf("Popping an empty stack should return an error.")
	}

	stackLength := 10

	for i := 1; i <= stackLength; i++ {
		stack.Push(i)
	}

	for i := stackLength; i >= 1; i-- {
		item, err := stack.Pop()

		if err != nil {
			t.Errorf("Popping a populated stack should not return an error")
		}

		if item != i {
			t.Errorf("Error popping element from stack. Got %v, expected %v", item, i)
		}

		currentLength := stack.Length()

		if currentLength != i-1 {
			t.Errorf("Popping the top element from stack should decrease the length. Got %v, expected %v", currentLength, i-1)
		}
	}
}
