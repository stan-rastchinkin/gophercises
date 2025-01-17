package queue

import (
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

		if queue.Len() != testElQty {
			t.Errorf("Incorrect length %d\n", queue.Len())
		}

		element := queue.Pull()

		if element != 1 {
			t.Errorf("The element must be %d\n", 1)
		}
	})

	t.Run("Pulling from empty queue returns zero value", func(t *testing.T) {
		intQueue := New[int]()

		zeroEl1 := intQueue.Pull()

		if zeroEl1 != 0 {
			t.Errorf("zeroEl1 must be 0\n")
		}

		type st struct{}

		stQueue := New[*st]()

		zeroEl2 := stQueue.Pull()

		if zeroEl2 != nil {
			t.Errorf("zeroEl2 must be nil\n")
		}
	})
}
