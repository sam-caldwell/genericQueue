package GenericQueue

import "testing"

func TestQueue_struct(t *testing.T) {
	q := Queue[int]{
		head:  &(QueueNode[int]{}),
		tail:  &(QueueNode[int]{}),
		count: 0,
	}
	q.lock.Lock()
	q.lock.Unlock()
}
