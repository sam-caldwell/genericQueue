package GenericQueue

import (
	"sync"
)

type Queue[T any] struct {
	lock  sync.Mutex
	head  *QueueNode[T]
	tail  *QueueNode[T]
	count uint
}
