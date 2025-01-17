package queue

import "testing"

func TestQueue(t *testing.T) {
	testElQty := 3

	t.Run("Test queue", func(t *testing.T) {
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
}
