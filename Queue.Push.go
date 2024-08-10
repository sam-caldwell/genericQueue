package GenericQueue

// Push - Add a new element to the tail-end of the GenericQueue
func (queue *Queue[T]) Push(data T) {

	queue.lock.Lock()
	defer queue.lock.Unlock()

	newNode := &QueueNode[T]{data: data}

	if queue.tail == nil { // List is empty
		queue.head = newNode
		queue.tail = newNode
	} else {
		queue.tail.nextPtr = newNode
		queue.tail = newNode
	}

	queue.count++

}
