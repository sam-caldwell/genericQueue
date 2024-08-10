package genericQueue

import "fmt"

// Pop - Remove and return the item at the head of the GenericQueue
func (queue *Queue[T]) Pop() (data T, err error) {

	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.head == nil {
		var nothing T
		return nothing, fmt.Errorf("empty queue")
	}

	currentRecord := queue.head
	if currentRecord.nextPtr == nil {
		queue.head = nil
	}
	//move head to next node.
	queue.head = currentRecord.nextPtr
	currentRecord.nextPtr = nil

	queue.count-- // Decrement count

	return currentRecord.data, err

}
