package queue

import (
	"container/list"
)

// 4. But that's not the case:
// our interface strictly defines what will be available
type Queue[T any] interface {
	Pull() T
	Push(T)
	Len() int
	Peek() []T
	// Purge() -- TODO
}

// 1. This struct directly extends a List
type queueState[T any] struct {
	*list.List
}

func (qs *queueState[T]) Pull() T {
	el := qs.Front()
	if el == nil {
		var zero T

		return zero
	}
	val := qs.Remove(el).(T) // type assertion

	return val
}

func (qs *queueState[T]) Push(element T) {
	qs.PushBack(element)
}

// 2. We do not need to define Len() method
// becuase it is already present on List
// and queueState is effectively a List

// func (qs *queueState) Len() int {
// 	return qs.Len()
// }

func (qs queueState[T]) Peek() []T {
	var items []T
	currentEl := qs.Front()

	for currentEl != nil {
		items = append(items, currentEl.Value.(T))
		currentEl = currentEl.Next()
	}

	return items
}

// 3. Then we might think that all the methods
// of List will be awailable on Queue
func New[T any]() Queue[T] {
	return &queueState[T]{list.New()}
}
