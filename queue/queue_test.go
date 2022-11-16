package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	queue := NewQueue()

	if !queue.IsEmpty() {
		t.Errorf("A newly instantiated queue should have no elements. Got %v, expected %v", queue.Length(), 0)
	}

	if queue.Length() != 0 {
		t.Errorf("A newly instantiated queue should have no elements. Got %v, expected %v", queue.Length(), 0)
	}
}

func TestQueuePush(t *testing.T) {
	queue := NewQueue()
	maxLength := 10

	for i := 1; i <= maxLength; i++ {
		queue.Add(i)

		if queue.Last.Data != i {
			t.Errorf("Error pushing element to queue. Got %v, expected %v on the 'last' element", queue.Last.Data, i)
		}
	}

	if queue.First.Data != 1 {
		t.Errorf("Error pushing first element to queue. Got %v, expected %v on the 'tail' element", queue.First.Data, 1)
	}
}

func TestQueueLength(t *testing.T) {
	queue := NewQueue()
	queueLength := 10

	for i := 1; i <= queueLength; i++ {
		queue.Add(i)
	}

	if queue.Length() != queueLength {
		t.Errorf("Wrong queue length. Got %v, expected %v", queue.Length(), queueLength)
	}
}

func TestQueueIsEmpty(t *testing.T) {
	queue := NewQueue()

	if !queue.IsEmpty() {
		t.Errorf("Queue should be empty. Got %v, expected %v", queue.Length(), 0)
	}

	queue.Add(1)

	if queue.IsEmpty() {
		t.Errorf("Queue should not be empty. Got %v, expected %v", queue.Length(), 0)
	}
}

func TestQueuePeek(t *testing.T) {
	queue := NewQueue()

	_, err := queue.Peek()

	if err == nil {
		t.Errorf("Peeking an empty queue should return an error.")
	}

	queueLength := 10

	for i := 1; i <= queueLength; i++ {
		queue.Add(i)

		item, err := queue.Peek()

		if err != nil {
			t.Errorf("Peeking a populated queue should not return an error")
		}

		if item != 1 {
			t.Errorf("Error peeking element from queue. Got %v, expected %v", item, i)
		}
	}

	currentLength := queue.Length()

	if currentLength != queueLength {
		t.Errorf("Peeking the top element from queue should not decrease the length. Got %v, expected %v", currentLength, queueLength)
	}
}

func TestQueuePop(t *testing.T) {
	queue := NewQueue()

	_, err := queue.Remove()

	if err == nil {
		t.Errorf("Removing an element from an empty queue should return an error.")
	}

	queueLength := 10

	for i := 1; i <= queueLength; i++ {
		queue.Add(i)
	}

	for i := 1; i <= queueLength; i++ {
		item, err := queue.Remove()

		if err != nil {
			t.Errorf("Removing an element from a populated queue should not return an error")
		}

		if item != i {
			t.Errorf("Error removing element from queue. Got %v, expected %v", item, i)
		}

		currentLength := queue.Length()

		if currentLength != queueLength-i {
			t.Errorf("Popping the top element from queue should decrease the length. Got %v, expected %v", currentLength, queueLength-1)
		}
	}

	if queue.First != nil {
		t.Errorf("An empty queue should have a nil first element. Got %v, expected %v", queue.First, "nil")
	}

	if queue.Last != nil {
		t.Errorf("An empty queue should have a nil Last element. Got %v, expected %v", queue.Last, "nil")
	}

	if queue.Length() != 0 {
		t.Errorf("An empty queue should have a valid length. Got %v, expected %v", queue.Length(), 0)
	}

	if !queue.IsEmpty() {
		t.Errorf("An empty queue should have be empty. Got %v, expected %v", queue.IsEmpty(), true)
	}
}
