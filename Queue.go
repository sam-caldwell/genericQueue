package GenericQueue

import (
	"sync"
)

// Queue - A queue based on a linked list
type Queue[T any] struct {
	lock        sync.Mutex
	head        *QueueNode[T]
	tail        *QueueNode[T]
	count       uint
	compareFunc func(T, T) int
}
