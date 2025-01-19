package queue

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	testElQty := 3

	t.Run("Normal behavior", func(t *testing.T) {
		queue := New[int]()

		if queue.Len() != 0 {
			t.Errorf("Initialized with non zero length\n")
		}

		for i := 1; i <= testElQty; i++ {
			queue.Push(i)
		}

		items := queue.Peek()
		if !reflect.DeepEqual(items, []int{1, 2, 3}) {
			t.Errorf("Peek returned %v\n", items)
		}

		if queue.Len() != testElQty {
			t.Errorf("Incorrect length %d\n", queue.Len())
		}

		element, _ := queue.Pull()

		if element != 1 {
			t.Errorf("The element must be %d\n", 1)
		}
	})

	t.Run("Pulling from empty int queue", func(t *testing.T) {
		intQueue := New[int]()

		zeroEl, isEmpty := intQueue.Pull()

		if zeroEl != 0 {
			t.Errorf("zeroEl must be 0\n")
		}
		if !isEmpty {
			t.Errorf("isEmpty must be true\n")
		}
	})

	t.Run("Pulling from empty struct queue", func(t *testing.T) {
		type st struct{}

		stQueue := New[*st]()

		zeroEl, isEmpty := stQueue.Pull()

		if zeroEl != nil {
			t.Errorf("zeroEl must be nil\n")
		}
		if !isEmpty {
			t.Errorf("isEmpty must be true\n")
		}
	})
}
